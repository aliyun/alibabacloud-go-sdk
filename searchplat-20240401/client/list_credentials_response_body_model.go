// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListCredentialsResponseBody
	GetRequestId() *string
	SetResult(v []*ListCredentialsResponseBodyResult) *ListCredentialsResponseBody
	GetResult() []*ListCredentialsResponseBodyResult
	SetTotalCount(v int64) *ListCredentialsResponseBody
	GetTotalCount() *int64
}

type ListCredentialsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 33E4F0CA-F766-5803-B11C-70DC57A5A6E4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result []*ListCredentialsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCredentialsResponseBody) GetResult() []*ListCredentialsResponseBodyResult {
	return s.Result
}

func (s *ListCredentialsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCredentialsResponseBody) SetRequestId(v string) *ListCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCredentialsResponseBody) SetResult(v []*ListCredentialsResponseBodyResult) *ListCredentialsResponseBody {
	s.Result = v
	return s
}

func (s *ListCredentialsResponseBody) SetTotalCount(v int64) *ListCredentialsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCredentialsResponseBody) Validate() error {
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

type ListCredentialsResponseBodyResult struct {
	// The workspace ID.
	//
	// example:
	//
	// 123123
	AppGroupId *int64 `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// Indicates whether the credential is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The access credential token.
	//
	// example:
	//
	// OS-****
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// The credential type.
	//
	// - api-token.
	//
	// example:
	//
	// api-token
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListCredentialsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponseBodyResult) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *ListCredentialsResponseBodyResult) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListCredentialsResponseBodyResult) GetToken() *string {
	return s.Token
}

func (s *ListCredentialsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListCredentialsResponseBodyResult) SetAppGroupId(v int64) *ListCredentialsResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *ListCredentialsResponseBodyResult) SetEnabled(v bool) *ListCredentialsResponseBodyResult {
	s.Enabled = &v
	return s
}

func (s *ListCredentialsResponseBodyResult) SetToken(v string) *ListCredentialsResponseBodyResult {
	s.Token = &v
	return s
}

func (s *ListCredentialsResponseBodyResult) SetType(v string) *ListCredentialsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListCredentialsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
