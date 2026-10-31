// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyResourcePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyResourcePermissionResponseBody
	GetCode() *string
	SetData(v *ApplyResourcePermissionResponseBodyData) *ApplyResourcePermissionResponseBody
	GetData() *ApplyResourcePermissionResponseBodyData
	SetHttpStatusCode(v int32) *ApplyResourcePermissionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ApplyResourcePermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApplyResourcePermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyResourcePermissionResponseBody
	GetSuccess() *bool
}

type ApplyResourcePermissionResponseBody struct {
	// The error code of the request. A value of OK indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The application result.
	Data *ApplyResourcePermissionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message of the request.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyResourcePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourcePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyResourcePermissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyResourcePermissionResponseBody) GetData() *ApplyResourcePermissionResponseBodyData {
	return s.Data
}

func (s *ApplyResourcePermissionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ApplyResourcePermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyResourcePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyResourcePermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyResourcePermissionResponseBody) SetCode(v string) *ApplyResourcePermissionResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyResourcePermissionResponseBody) SetData(v *ApplyResourcePermissionResponseBodyData) *ApplyResourcePermissionResponseBody {
	s.Data = v
	return s
}

func (s *ApplyResourcePermissionResponseBody) SetHttpStatusCode(v int32) *ApplyResourcePermissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ApplyResourcePermissionResponseBody) SetMessage(v string) *ApplyResourcePermissionResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyResourcePermissionResponseBody) SetRequestId(v string) *ApplyResourcePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyResourcePermissionResponseBody) SetSuccess(v bool) *ApplyResourcePermissionResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyResourcePermissionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyResourcePermissionResponseBodyData struct {
	// The list of application IDs.
	ApplyIds []*string `json:"ApplyIds,omitempty" xml:"ApplyIds,omitempty" type:"Repeated"`
	// The list of approval IDs.
	BpmsIds []*string `json:"BpmsIds,omitempty" xml:"BpmsIds,omitempty" type:"Repeated"`
}

func (s ApplyResourcePermissionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourcePermissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ApplyResourcePermissionResponseBodyData) GetApplyIds() []*string {
	return s.ApplyIds
}

func (s *ApplyResourcePermissionResponseBodyData) GetBpmsIds() []*string {
	return s.BpmsIds
}

func (s *ApplyResourcePermissionResponseBodyData) SetApplyIds(v []*string) *ApplyResourcePermissionResponseBodyData {
	s.ApplyIds = v
	return s
}

func (s *ApplyResourcePermissionResponseBodyData) SetBpmsIds(v []*string) *ApplyResourcePermissionResponseBodyData {
	s.BpmsIds = v
	return s
}

func (s *ApplyResourcePermissionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
