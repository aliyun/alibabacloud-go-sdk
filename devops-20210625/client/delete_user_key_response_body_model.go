// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteUserKeyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteUserKeyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteUserKeyResponseBody
	GetRequestId() *string
	SetResult(v *DeleteUserKeyResponseBodyResult) *DeleteUserKeyResponseBody
	GetResult() *DeleteUserKeyResponseBodyResult
	SetSuccess(v bool) *DeleteUserKeyResponseBody
	GetSuccess() *bool
}

type DeleteUserKeyResponseBody struct {
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
	Result    *DeleteUserKeyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteUserKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserKeyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteUserKeyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteUserKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserKeyResponseBody) GetResult() *DeleteUserKeyResponseBodyResult {
	return s.Result
}

func (s *DeleteUserKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUserKeyResponseBody) SetErrorCode(v string) *DeleteUserKeyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteUserKeyResponseBody) SetErrorMessage(v string) *DeleteUserKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteUserKeyResponseBody) SetRequestId(v string) *DeleteUserKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserKeyResponseBody) SetResult(v *DeleteUserKeyResponseBodyResult) *DeleteUserKeyResponseBody {
	s.Result = v
	return s
}

func (s *DeleteUserKeyResponseBody) SetSuccess(v bool) *DeleteUserKeyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUserKeyResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteUserKeyResponseBodyResult struct {
	// example:
	//
	// xxx
	Context *string `json:"context,omitempty" xml:"context,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// xxx
	FingerPrint *string `json:"fingerPrint,omitempty" xml:"fingerPrint,omitempty"`
	// example:
	//
	// 50998
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// ALL
	KeyScope *string `json:"keyScope,omitempty" xml:"keyScope,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	LastUsedTime *string `json:"lastUsedTime,omitempty" xml:"lastUsedTime,omitempty"`
	// example:
	//
	// xxx
	PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty"`
	// example:
	//
	// xxx
	ShaContext *string `json:"shaContext,omitempty" xml:"shaContext,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s DeleteUserKeyResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserKeyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteUserKeyResponseBodyResult) GetContext() *string {
	return s.Context
}

func (s *DeleteUserKeyResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DeleteUserKeyResponseBodyResult) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DeleteUserKeyResponseBodyResult) GetFingerPrint() *string {
	return s.FingerPrint
}

func (s *DeleteUserKeyResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *DeleteUserKeyResponseBodyResult) GetKeyScope() *string {
	return s.KeyScope
}

func (s *DeleteUserKeyResponseBodyResult) GetLastUsedTime() *string {
	return s.LastUsedTime
}

func (s *DeleteUserKeyResponseBodyResult) GetPublicKey() *string {
	return s.PublicKey
}

func (s *DeleteUserKeyResponseBodyResult) GetShaContext() *string {
	return s.ShaContext
}

func (s *DeleteUserKeyResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *DeleteUserKeyResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DeleteUserKeyResponseBodyResult) SetContext(v string) *DeleteUserKeyResponseBodyResult {
	s.Context = &v
	return s
}

func (s *DeleteUserKeyResponseBodyResult) SetCreatedAt(v string) *DeleteUserKeyResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *DeleteUserKeyResponseBodyResult) SetExpireTime(v string) *DeleteUserKeyResponseBodyResult {
	s.ExpireTime = &v
	return s
}

func (s *DeleteUserKeyResponseBodyResult) SetFingerPrint(v string) *DeleteUserKeyResponseBodyResult {
	s.FingerPrint = &v
	return s
}

func (s *DeleteUserKeyResponseBodyResult) SetId(v int64) *DeleteUserKeyResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DeleteUserKeyResponseBodyResult) SetKeyScope(v string) *DeleteUserKeyResponseBodyResult {
	s.KeyScope = &v
	return s
}

func (s *DeleteUserKeyResponseBodyResult) SetLastUsedTime(v string) *DeleteUserKeyResponseBodyResult {
	s.LastUsedTime = &v
	return s
}

func (s *DeleteUserKeyResponseBodyResult) SetPublicKey(v string) *DeleteUserKeyResponseBodyResult {
	s.PublicKey = &v
	return s
}

func (s *DeleteUserKeyResponseBodyResult) SetShaContext(v string) *DeleteUserKeyResponseBodyResult {
	s.ShaContext = &v
	return s
}

func (s *DeleteUserKeyResponseBodyResult) SetTitle(v string) *DeleteUserKeyResponseBodyResult {
	s.Title = &v
	return s
}

func (s *DeleteUserKeyResponseBodyResult) SetUpdatedAt(v string) *DeleteUserKeyResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *DeleteUserKeyResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
