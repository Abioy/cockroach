// Copyright 2015 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Peter Mattis (peter@cockroachlabs.com)

package sqlwire

import "strconv"

func (d Datum) String() string {
	if d.BoolVal != nil {
		if *d.BoolVal {
			return "true"
		}
		return "false"
	}
	if d.IntVal != nil {
		return strconv.FormatInt(*d.IntVal, 10)
	}
	if d.FloatVal != nil {
		return strconv.FormatFloat(*d.FloatVal, 'g', -1, 64)
	}
	if d.BytesVal != nil {
		return string(d.BytesVal)
	}
	if d.StringVal != nil {
		return *d.StringVal
	}
	return "NULL"
}
