// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDoctorApplicationResponseBodyData) *GetDoctorApplicationResponseBody
	GetData() *GetDoctorApplicationResponseBodyData
}

type GetDoctorApplicationResponseBody struct {
	// The data returned.
	Data *GetDoctorApplicationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s GetDoctorApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationResponseBody) GetData() *GetDoctorApplicationResponseBodyData {
	return s.Data
}

func (s *GetDoctorApplicationResponseBody) SetData(v *GetDoctorApplicationResponseBodyData) *GetDoctorApplicationResponseBody {
	s.Data = v
	return s
}

func (s *GetDoctorApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDoctorApplicationResponseBodyData struct {
	// The diagnostics list.
	Suggestions []*string `json:"suggestions,omitempty" xml:"suggestions,omitempty" type:"Repeated"`
}

func (s GetDoctorApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationResponseBodyData) GetSuggestions() []*string {
	return s.Suggestions
}

func (s *GetDoctorApplicationResponseBodyData) SetSuggestions(v []*string) *GetDoctorApplicationResponseBodyData {
	s.Suggestions = v
	return s
}

func (s *GetDoctorApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
