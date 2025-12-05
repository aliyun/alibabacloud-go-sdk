// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllRegions(v map[string]*string) *GetAllRegionsResponseBody
	GetAllRegions() map[string]*string
	SetCode(v string) *GetAllRegionsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetAllRegionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAllRegionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAllRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAllRegionsResponseBody
	GetSuccess() *bool
}

type GetAllRegionsResponseBody struct {
	// The supported regions.
	AllRegions map[string]*string `json:"AllRegions,omitempty" xml:"AllRegions,omitempty"`
	// The system status code. If the request was successful, no data is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request was successful, no data is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 73D16B8D-0FCD-5596-B7BE-A47042989318
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAllRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAllRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllRegionsResponseBody) GetAllRegions() map[string]*string {
	return s.AllRegions
}

func (s *GetAllRegionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAllRegionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAllRegionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAllRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAllRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAllRegionsResponseBody) SetAllRegions(v map[string]*string) *GetAllRegionsResponseBody {
	s.AllRegions = v
	return s
}

func (s *GetAllRegionsResponseBody) SetCode(v string) *GetAllRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetAllRegionsResponseBody) SetHttpStatusCode(v int32) *GetAllRegionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAllRegionsResponseBody) SetMessage(v string) *GetAllRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetAllRegionsResponseBody) SetRequestId(v string) *GetAllRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllRegionsResponseBody) SetSuccess(v bool) *GetAllRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *GetAllRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}
