// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryApiKeysResponseBody
	GetCode() *string
	SetData(v []*QueryApiKeysResponseBodyData) *QueryApiKeysResponseBody
	GetData() []*QueryApiKeysResponseBodyData
	SetMessage(v string) *QueryApiKeysResponseBody
	GetMessage() *string
	SetRetryAble(v bool) *QueryApiKeysResponseBody
	GetRetryAble() *bool
	SetSuccess(v bool) *QueryApiKeysResponseBody
	GetSuccess() *bool
}

type QueryApiKeysResponseBody struct {
	// example:
	//
	// 0
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data []*QueryApiKeysResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// false
	RetryAble *bool `json:"retryAble,omitempty" xml:"retryAble,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *QueryApiKeysResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryApiKeysResponseBody) GetData() []*QueryApiKeysResponseBodyData {
	return s.Data
}

func (s *QueryApiKeysResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryApiKeysResponseBody) GetRetryAble() *bool {
	return s.RetryAble
}

func (s *QueryApiKeysResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryApiKeysResponseBody) SetCode(v string) *QueryApiKeysResponseBody {
	s.Code = &v
	return s
}

func (s *QueryApiKeysResponseBody) SetData(v []*QueryApiKeysResponseBodyData) *QueryApiKeysResponseBody {
	s.Data = v
	return s
}

func (s *QueryApiKeysResponseBody) SetMessage(v string) *QueryApiKeysResponseBody {
	s.Message = &v
	return s
}

func (s *QueryApiKeysResponseBody) SetRetryAble(v bool) *QueryApiKeysResponseBody {
	s.RetryAble = &v
	return s
}

func (s *QueryApiKeysResponseBody) SetSuccess(v bool) *QueryApiKeysResponseBody {
	s.Success = &v
	return s
}

func (s *QueryApiKeysResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryApiKeysResponseBodyData struct {
	// example:
	//
	// 2024-12-31T23:59:59Z
	ExpiresAt *string `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// a1b2c3d4e5f6...
	KeyHash *string `json:"keyHash,omitempty" xml:"keyHash,omitempty"`
	// **API Key ID**
	//
	// example:
	//
	// key_001
	KeyId *string `json:"keyId,omitempty" xml:"keyId,omitempty"`
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 100
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryApiKeysResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryApiKeysResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryApiKeysResponseBodyData) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *QueryApiKeysResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryApiKeysResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryApiKeysResponseBodyData) GetKeyHash() *string {
	return s.KeyHash
}

func (s *QueryApiKeysResponseBodyData) GetKeyId() *string {
	return s.KeyId
}

func (s *QueryApiKeysResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryApiKeysResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryApiKeysResponseBodyData) SetExpiresAt(v string) *QueryApiKeysResponseBodyData {
	s.ExpiresAt = &v
	return s
}

func (s *QueryApiKeysResponseBodyData) SetGmtCreate(v string) *QueryApiKeysResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryApiKeysResponseBodyData) SetGmtModified(v string) *QueryApiKeysResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryApiKeysResponseBodyData) SetKeyHash(v string) *QueryApiKeysResponseBodyData {
	s.KeyHash = &v
	return s
}

func (s *QueryApiKeysResponseBodyData) SetKeyId(v string) *QueryApiKeysResponseBodyData {
	s.KeyId = &v
	return s
}

func (s *QueryApiKeysResponseBodyData) SetStatus(v string) *QueryApiKeysResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryApiKeysResponseBodyData) SetTenantId(v string) *QueryApiKeysResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *QueryApiKeysResponseBodyData) Validate() error {
	return dara.Validate(s)
}
