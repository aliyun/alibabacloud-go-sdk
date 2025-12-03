// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListUserKeysResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListUserKeysResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListUserKeysResponseBody
	GetRequestId() *string
	SetResult(v []*ListUserKeysResponseBodyResult) *ListUserKeysResponseBody
	GetResult() []*ListUserKeysResponseBodyResult
	SetSuccess(v bool) *ListUserKeysResponseBody
	GetSuccess() *bool
}

type ListUserKeysResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Openapi.RequestError
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// F7B85D1B-D1C2-140F-A039-341859F130B9
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListUserKeysResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListUserKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserKeysResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListUserKeysResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListUserKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserKeysResponseBody) GetResult() []*ListUserKeysResponseBodyResult {
	return s.Result
}

func (s *ListUserKeysResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserKeysResponseBody) SetErrorCode(v string) *ListUserKeysResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUserKeysResponseBody) SetErrorMessage(v string) *ListUserKeysResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListUserKeysResponseBody) SetRequestId(v string) *ListUserKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserKeysResponseBody) SetResult(v []*ListUserKeysResponseBodyResult) *ListUserKeysResponseBody {
	s.Result = v
	return s
}

func (s *ListUserKeysResponseBody) SetSuccess(v bool) *ListUserKeysResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserKeysResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserKeysResponseBodyResult struct {
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
	// 5240
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
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListUserKeysResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListUserKeysResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListUserKeysResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListUserKeysResponseBodyResult) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListUserKeysResponseBodyResult) GetFingerPrint() *string {
	return s.FingerPrint
}

func (s *ListUserKeysResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListUserKeysResponseBodyResult) GetKeyScope() *string {
	return s.KeyScope
}

func (s *ListUserKeysResponseBodyResult) GetLastUsedTime() *string {
	return s.LastUsedTime
}

func (s *ListUserKeysResponseBodyResult) GetPublicKey() *string {
	return s.PublicKey
}

func (s *ListUserKeysResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *ListUserKeysResponseBodyResult) SetCreatedAt(v string) *ListUserKeysResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListUserKeysResponseBodyResult) SetExpireTime(v string) *ListUserKeysResponseBodyResult {
	s.ExpireTime = &v
	return s
}

func (s *ListUserKeysResponseBodyResult) SetFingerPrint(v string) *ListUserKeysResponseBodyResult {
	s.FingerPrint = &v
	return s
}

func (s *ListUserKeysResponseBodyResult) SetId(v int64) *ListUserKeysResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListUserKeysResponseBodyResult) SetKeyScope(v string) *ListUserKeysResponseBodyResult {
	s.KeyScope = &v
	return s
}

func (s *ListUserKeysResponseBodyResult) SetLastUsedTime(v string) *ListUserKeysResponseBodyResult {
	s.LastUsedTime = &v
	return s
}

func (s *ListUserKeysResponseBodyResult) SetPublicKey(v string) *ListUserKeysResponseBodyResult {
	s.PublicKey = &v
	return s
}

func (s *ListUserKeysResponseBodyResult) SetTitle(v string) *ListUserKeysResponseBodyResult {
	s.Title = &v
	return s
}

func (s *ListUserKeysResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
