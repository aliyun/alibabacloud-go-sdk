// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSshKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ResetSshKeyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ResetSshKeyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ResetSshKeyResponseBody
	GetRequestId() *string
	SetSshKey(v *ResetSshKeyResponseBodySshKey) *ResetSshKeyResponseBody
	GetSshKey() *ResetSshKeyResponseBodySshKey
	SetSuccess(v bool) *ResetSshKeyResponseBody
	GetSuccess() *bool
}

type ResetSshKeyResponseBody struct {
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
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SshKey    *ResetSshKeyResponseBodySshKey `json:"sshKey,omitempty" xml:"sshKey,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ResetSshKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetSshKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ResetSshKeyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ResetSshKeyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ResetSshKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetSshKeyResponseBody) GetSshKey() *ResetSshKeyResponseBodySshKey {
	return s.SshKey
}

func (s *ResetSshKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResetSshKeyResponseBody) SetErrorCode(v string) *ResetSshKeyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResetSshKeyResponseBody) SetErrorMessage(v string) *ResetSshKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResetSshKeyResponseBody) SetRequestId(v string) *ResetSshKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetSshKeyResponseBody) SetSshKey(v *ResetSshKeyResponseBodySshKey) *ResetSshKeyResponseBody {
	s.SshKey = v
	return s
}

func (s *ResetSshKeyResponseBody) SetSuccess(v bool) *ResetSshKeyResponseBody {
	s.Success = &v
	return s
}

func (s *ResetSshKeyResponseBody) Validate() error {
	if s.SshKey != nil {
		if err := s.SshKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetSshKeyResponseBodySshKey struct {
	// example:
	//
	// 1212
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// saaaaaaaaaaaaaaaa
	PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty"`
}

func (s ResetSshKeyResponseBodySshKey) String() string {
	return dara.Prettify(s)
}

func (s ResetSshKeyResponseBodySshKey) GoString() string {
	return s.String()
}

func (s *ResetSshKeyResponseBodySshKey) GetId() *int64 {
	return s.Id
}

func (s *ResetSshKeyResponseBodySshKey) GetPublicKey() *string {
	return s.PublicKey
}

func (s *ResetSshKeyResponseBodySshKey) SetId(v int64) *ResetSshKeyResponseBodySshKey {
	s.Id = &v
	return s
}

func (s *ResetSshKeyResponseBodySshKey) SetPublicKey(v string) *ResetSshKeyResponseBodySshKey {
	s.PublicKey = &v
	return s
}

func (s *ResetSshKeyResponseBodySshKey) Validate() error {
	return dara.Validate(s)
}
