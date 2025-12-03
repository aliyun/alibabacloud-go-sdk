// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSshKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateSshKeyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateSshKeyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateSshKeyResponseBody
	GetRequestId() *string
	SetSshKey(v *CreateSshKeyResponseBodySshKey) *CreateSshKeyResponseBody
	GetSshKey() *CreateSshKeyResponseBodySshKey
	SetSuccess(v bool) *CreateSshKeyResponseBody
	GetSuccess() *bool
}

type CreateSshKeyResponseBody struct {
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
	RequestId *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SshKey    *CreateSshKeyResponseBodySshKey `json:"sshKey,omitempty" xml:"sshKey,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateSshKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSshKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSshKeyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateSshKeyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateSshKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSshKeyResponseBody) GetSshKey() *CreateSshKeyResponseBodySshKey {
	return s.SshKey
}

func (s *CreateSshKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSshKeyResponseBody) SetErrorCode(v string) *CreateSshKeyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSshKeyResponseBody) SetErrorMessage(v string) *CreateSshKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSshKeyResponseBody) SetRequestId(v string) *CreateSshKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSshKeyResponseBody) SetSshKey(v *CreateSshKeyResponseBodySshKey) *CreateSshKeyResponseBody {
	s.SshKey = v
	return s
}

func (s *CreateSshKeyResponseBody) SetSuccess(v bool) *CreateSshKeyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSshKeyResponseBody) Validate() error {
	if s.SshKey != nil {
		if err := s.SshKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSshKeyResponseBodySshKey struct {
	// example:
	//
	// 123
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// assssssssssss
	PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty"`
}

func (s CreateSshKeyResponseBodySshKey) String() string {
	return dara.Prettify(s)
}

func (s CreateSshKeyResponseBodySshKey) GoString() string {
	return s.String()
}

func (s *CreateSshKeyResponseBodySshKey) GetId() *int64 {
	return s.Id
}

func (s *CreateSshKeyResponseBodySshKey) GetPublicKey() *string {
	return s.PublicKey
}

func (s *CreateSshKeyResponseBodySshKey) SetId(v int64) *CreateSshKeyResponseBodySshKey {
	s.Id = &v
	return s
}

func (s *CreateSshKeyResponseBodySshKey) SetPublicKey(v string) *CreateSshKeyResponseBodySshKey {
	s.PublicKey = &v
	return s
}

func (s *CreateSshKeyResponseBodySshKey) Validate() error {
	return dara.Validate(s)
}
