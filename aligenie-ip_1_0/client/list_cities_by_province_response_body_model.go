// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCitiesByProvinceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ListCitiesByProvinceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCitiesByProvinceResponseBody
	GetRequestId() *string
	SetResult(v []*string) *ListCitiesByProvinceResponseBody
	GetResult() []*string
	SetStatusCode(v int32) *ListCitiesByProvinceResponseBody
	GetStatusCode() *int32
}

type ListCitiesByProvinceResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 860194F7-9593-50EA-8E53-BCEC0D325A00
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListCitiesByProvinceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCitiesByProvinceResponseBody) GoString() string {
	return s.String()
}

func (s *ListCitiesByProvinceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCitiesByProvinceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCitiesByProvinceResponseBody) GetResult() []*string {
	return s.Result
}

func (s *ListCitiesByProvinceResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCitiesByProvinceResponseBody) SetMessage(v string) *ListCitiesByProvinceResponseBody {
	s.Message = &v
	return s
}

func (s *ListCitiesByProvinceResponseBody) SetRequestId(v string) *ListCitiesByProvinceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCitiesByProvinceResponseBody) SetResult(v []*string) *ListCitiesByProvinceResponseBody {
	s.Result = v
	return s
}

func (s *ListCitiesByProvinceResponseBody) SetStatusCode(v int32) *ListCitiesByProvinceResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ListCitiesByProvinceResponseBody) Validate() error {
	return dara.Validate(s)
}
