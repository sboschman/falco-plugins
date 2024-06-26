// SPDX-License-Identifier: Apache-2.0
/*
Copyright (C) 2023 The Falco Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package registry

import (
	"bytes"
)

func encodeString(r encoder) string {
	buf := bytes.Buffer{}
	err := r.Encode(&buf)
	if err != nil {
		return "string encoding error: " + err.Error()
	}
	return buf.String()
}

// String implements the fmt.Stringer interface
func (r *SourcingCapability) String() string {
	return encodeString(r)
}

// String implements the fmt.Stringer interface
func (r *ExtractionCapability) String() string {
	return encodeString(r)
}

// String implements the fmt.Stringer interface
func (r *Capabilities) String() string {
	return encodeString(r)
}

// String implements the fmt.Stringer interface
func (r *Plugin) String() string {
	return encodeString(r)
}

// String implements the fmt.Stringer interface
func (r *Registry) String() string {
	return encodeString(r)
}
