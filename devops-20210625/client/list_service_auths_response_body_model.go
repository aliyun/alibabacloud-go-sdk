// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceAuthsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListServiceAuthsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListServiceAuthsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListServiceAuthsResponseBody
	GetRequestId() *string
	SetServiceAuths(v []*ListServiceAuthsResponseBodyServiceAuths) *ListServiceAuthsResponseBody
	GetServiceAuths() []*ListServiceAuthsResponseBodyServiceAuths
	SetSuccess(v bool) *ListServiceAuthsResponseBody
	GetSuccess() *bool
}

type ListServiceAuthsResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId    *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ServiceAuths []*ListServiceAuthsResponseBodyServiceAuths `json:"serviceAuths,omitempty" xml:"serviceAuths,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListServiceAuthsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceAuthsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceAuthsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListServiceAuthsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListServiceAuthsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceAuthsResponseBody) GetServiceAuths() []*ListServiceAuthsResponseBodyServiceAuths {
	return s.ServiceAuths
}

func (s *ListServiceAuthsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListServiceAuthsResponseBody) SetErrorCode(v string) *ListServiceAuthsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListServiceAuthsResponseBody) SetErrorMessage(v string) *ListServiceAuthsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListServiceAuthsResponseBody) SetRequestId(v string) *ListServiceAuthsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceAuthsResponseBody) SetServiceAuths(v []*ListServiceAuthsResponseBodyServiceAuths) *ListServiceAuthsResponseBody {
	s.ServiceAuths = v
	return s
}

func (s *ListServiceAuthsResponseBody) SetSuccess(v bool) *ListServiceAuthsResponseBody {
	s.Success = &v
	return s
}

func (s *ListServiceAuthsResponseBody) Validate() error {
	if s.ServiceAuths != nil {
		for _, item := range s.ServiceAuths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceAuthsResponseBodyServiceAuths struct {
	// example:
	//
	// 123
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 张三
	OwnerName *string `json:"ownerName,omitempty" xml:"ownerName,omitempty"`
	// example:
	//
	// 123456789
	OwnerStaffId *string `json:"ownerStaffId,omitempty" xml:"ownerStaffId,omitempty"`
	// example:
	//
	// Codeup
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListServiceAuthsResponseBodyServiceAuths) String() string {
	return dara.Prettify(s)
}

func (s ListServiceAuthsResponseBodyServiceAuths) GoString() string {
	return s.String()
}

func (s *ListServiceAuthsResponseBodyServiceAuths) GetId() *int64 {
	return s.Id
}

func (s *ListServiceAuthsResponseBodyServiceAuths) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListServiceAuthsResponseBodyServiceAuths) GetOwnerStaffId() *string {
	return s.OwnerStaffId
}

func (s *ListServiceAuthsResponseBodyServiceAuths) GetType() *string {
	return s.Type
}

func (s *ListServiceAuthsResponseBodyServiceAuths) SetId(v int64) *ListServiceAuthsResponseBodyServiceAuths {
	s.Id = &v
	return s
}

func (s *ListServiceAuthsResponseBodyServiceAuths) SetOwnerName(v string) *ListServiceAuthsResponseBodyServiceAuths {
	s.OwnerName = &v
	return s
}

func (s *ListServiceAuthsResponseBodyServiceAuths) SetOwnerStaffId(v string) *ListServiceAuthsResponseBodyServiceAuths {
	s.OwnerStaffId = &v
	return s
}

func (s *ListServiceAuthsResponseBodyServiceAuths) SetType(v string) *ListServiceAuthsResponseBodyServiceAuths {
	s.Type = &v
	return s
}

func (s *ListServiceAuthsResponseBodyServiceAuths) Validate() error {
	return dara.Validate(s)
}
