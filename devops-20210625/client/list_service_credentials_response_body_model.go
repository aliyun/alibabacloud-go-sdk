// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListServiceCredentialsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListServiceCredentialsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListServiceCredentialsResponseBody
	GetRequestId() *string
	SetServiceCredentials(v []*ListServiceCredentialsResponseBodyServiceCredentials) *ListServiceCredentialsResponseBody
	GetServiceCredentials() []*ListServiceCredentialsResponseBodyServiceCredentials
	SetSuccess(v bool) *ListServiceCredentialsResponseBody
	GetSuccess() *bool
}

type ListServiceCredentialsResponseBody struct {
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
	RequestId          *string                                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ServiceCredentials []*ListServiceCredentialsResponseBodyServiceCredentials `json:"serviceCredentials,omitempty" xml:"serviceCredentials,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListServiceCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceCredentialsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListServiceCredentialsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListServiceCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceCredentialsResponseBody) GetServiceCredentials() []*ListServiceCredentialsResponseBodyServiceCredentials {
	return s.ServiceCredentials
}

func (s *ListServiceCredentialsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListServiceCredentialsResponseBody) SetErrorCode(v string) *ListServiceCredentialsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListServiceCredentialsResponseBody) SetErrorMessage(v string) *ListServiceCredentialsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListServiceCredentialsResponseBody) SetRequestId(v string) *ListServiceCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceCredentialsResponseBody) SetServiceCredentials(v []*ListServiceCredentialsResponseBodyServiceCredentials) *ListServiceCredentialsResponseBody {
	s.ServiceCredentials = v
	return s
}

func (s *ListServiceCredentialsResponseBody) SetSuccess(v bool) *ListServiceCredentialsResponseBody {
	s.Success = &v
	return s
}

func (s *ListServiceCredentialsResponseBody) Validate() error {
	if s.ServiceCredentials != nil {
		for _, item := range s.ServiceCredentials {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceCredentialsResponseBodyServiceCredentials struct {
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

func (s ListServiceCredentialsResponseBodyServiceCredentials) String() string {
	return dara.Prettify(s)
}

func (s ListServiceCredentialsResponseBodyServiceCredentials) GoString() string {
	return s.String()
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) GetId() *int64 {
	return s.Id
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) GetOwnerStaffId() *string {
	return s.OwnerStaffId
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) GetType() *string {
	return s.Type
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) SetId(v int64) *ListServiceCredentialsResponseBodyServiceCredentials {
	s.Id = &v
	return s
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) SetOwnerName(v string) *ListServiceCredentialsResponseBodyServiceCredentials {
	s.OwnerName = &v
	return s
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) SetOwnerStaffId(v string) *ListServiceCredentialsResponseBodyServiceCredentials {
	s.OwnerStaffId = &v
	return s
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) SetType(v string) *ListServiceCredentialsResponseBodyServiceCredentials {
	s.Type = &v
	return s
}

func (s *ListServiceCredentialsResponseBodyServiceCredentials) Validate() error {
	return dara.Validate(s)
}
