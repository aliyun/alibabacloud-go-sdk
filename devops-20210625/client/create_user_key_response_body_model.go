// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateUserKeyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateUserKeyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateUserKeyResponseBody
	GetRequestId() *string
	SetResult(v *CreateUserKeyResponseBodyResult) *CreateUserKeyResponseBody
	GetResult() *CreateUserKeyResponseBodyResult
	SetSuccess(v bool) *CreateUserKeyResponseBody
	GetSuccess() *bool
}

type CreateUserKeyResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateUserKeyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateUserKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserKeyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateUserKeyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateUserKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserKeyResponseBody) GetResult() *CreateUserKeyResponseBodyResult {
	return s.Result
}

func (s *CreateUserKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateUserKeyResponseBody) SetErrorCode(v string) *CreateUserKeyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateUserKeyResponseBody) SetErrorMessage(v string) *CreateUserKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateUserKeyResponseBody) SetRequestId(v string) *CreateUserKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserKeyResponseBody) SetResult(v *CreateUserKeyResponseBodyResult) *CreateUserKeyResponseBody {
	s.Result = v
	return s
}

func (s *CreateUserKeyResponseBody) SetSuccess(v bool) *CreateUserKeyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUserKeyResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserKeyResponseBodyResult struct {
	// example:
	//
	// 2022-03-12 12:00:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 2022-03-12 12:00:00
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// xxx
	FingerPrint *string `json:"fingerPrint,omitempty" xml:"fingerPrint,omitempty"`
	// example:
	//
	// 11072
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// ALL
	KeyScope *string `json:"keyScope,omitempty" xml:"keyScope,omitempty"`
	// example:
	//
	// 2022-03-12 12:00:00
	LastUsedTime *string `json:"lastUsedTime,omitempty" xml:"lastUsedTime,omitempty"`
	// example:
	//
	// xxx
	PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty"`
	// example:
	//
	// My Title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateUserKeyResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateUserKeyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateUserKeyResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateUserKeyResponseBodyResult) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateUserKeyResponseBodyResult) GetFingerPrint() *string {
	return s.FingerPrint
}

func (s *CreateUserKeyResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *CreateUserKeyResponseBodyResult) GetKeyScope() *string {
	return s.KeyScope
}

func (s *CreateUserKeyResponseBodyResult) GetLastUsedTime() *string {
	return s.LastUsedTime
}

func (s *CreateUserKeyResponseBodyResult) GetPublicKey() *string {
	return s.PublicKey
}

func (s *CreateUserKeyResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *CreateUserKeyResponseBodyResult) SetCreatedAt(v string) *CreateUserKeyResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateUserKeyResponseBodyResult) SetExpireTime(v string) *CreateUserKeyResponseBodyResult {
	s.ExpireTime = &v
	return s
}

func (s *CreateUserKeyResponseBodyResult) SetFingerPrint(v string) *CreateUserKeyResponseBodyResult {
	s.FingerPrint = &v
	return s
}

func (s *CreateUserKeyResponseBodyResult) SetId(v int64) *CreateUserKeyResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateUserKeyResponseBodyResult) SetKeyScope(v string) *CreateUserKeyResponseBodyResult {
	s.KeyScope = &v
	return s
}

func (s *CreateUserKeyResponseBodyResult) SetLastUsedTime(v string) *CreateUserKeyResponseBodyResult {
	s.LastUsedTime = &v
	return s
}

func (s *CreateUserKeyResponseBodyResult) SetPublicKey(v string) *CreateUserKeyResponseBodyResult {
	s.PublicKey = &v
	return s
}

func (s *CreateUserKeyResponseBodyResult) SetTitle(v string) *CreateUserKeyResponseBodyResult {
	s.Title = &v
	return s
}

func (s *CreateUserKeyResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
