// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUpgradeReleaseVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListUpgradeReleaseVersionsResponseBodyData) *ListUpgradeReleaseVersionsResponseBody
	GetData() []*ListUpgradeReleaseVersionsResponseBodyData
	SetErrorCode(v string) *ListUpgradeReleaseVersionsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListUpgradeReleaseVersionsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *ListUpgradeReleaseVersionsResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *ListUpgradeReleaseVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUpgradeReleaseVersionsResponseBody
	GetSuccess() *bool
}

type ListUpgradeReleaseVersionsResponseBody struct {
	Data []*ListUpgradeReleaseVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F6DFB9EA-0E2A-52AC-BFD5-1ADBBFFB6A2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUpgradeReleaseVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUpgradeReleaseVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUpgradeReleaseVersionsResponseBody) GetData() []*ListUpgradeReleaseVersionsResponseBodyData {
	return s.Data
}

func (s *ListUpgradeReleaseVersionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListUpgradeReleaseVersionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListUpgradeReleaseVersionsResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ListUpgradeReleaseVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUpgradeReleaseVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUpgradeReleaseVersionsResponseBody) SetData(v []*ListUpgradeReleaseVersionsResponseBodyData) *ListUpgradeReleaseVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListUpgradeReleaseVersionsResponseBody) SetErrorCode(v string) *ListUpgradeReleaseVersionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUpgradeReleaseVersionsResponseBody) SetErrorMessage(v string) *ListUpgradeReleaseVersionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListUpgradeReleaseVersionsResponseBody) SetHttpStatusCode(v string) *ListUpgradeReleaseVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUpgradeReleaseVersionsResponseBody) SetRequestId(v string) *ListUpgradeReleaseVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUpgradeReleaseVersionsResponseBody) SetSuccess(v bool) *ListUpgradeReleaseVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListUpgradeReleaseVersionsResponseBody) Validate() error {
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

type ListUpgradeReleaseVersionsResponseBodyData struct {
	// example:
	//
	// https://help.aliyun.com/zh/hologres/product-overview/release-notes
	ReleaseNotesUrl *string `json:"ReleaseNotesUrl,omitempty" xml:"ReleaseNotesUrl,omitempty"`
	// example:
	//
	// Stable
	Type            *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	UpgradeStrategy []*string `json:"UpgradeStrategy,omitempty" xml:"UpgradeStrategy,omitempty" type:"Repeated"`
	// example:
	//
	// r2.2.47
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListUpgradeReleaseVersionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUpgradeReleaseVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUpgradeReleaseVersionsResponseBodyData) GetReleaseNotesUrl() *string {
	return s.ReleaseNotesUrl
}

func (s *ListUpgradeReleaseVersionsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListUpgradeReleaseVersionsResponseBodyData) GetUpgradeStrategy() []*string {
	return s.UpgradeStrategy
}

func (s *ListUpgradeReleaseVersionsResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ListUpgradeReleaseVersionsResponseBodyData) SetReleaseNotesUrl(v string) *ListUpgradeReleaseVersionsResponseBodyData {
	s.ReleaseNotesUrl = &v
	return s
}

func (s *ListUpgradeReleaseVersionsResponseBodyData) SetType(v string) *ListUpgradeReleaseVersionsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListUpgradeReleaseVersionsResponseBodyData) SetUpgradeStrategy(v []*string) *ListUpgradeReleaseVersionsResponseBodyData {
	s.UpgradeStrategy = v
	return s
}

func (s *ListUpgradeReleaseVersionsResponseBodyData) SetVersion(v string) *ListUpgradeReleaseVersionsResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListUpgradeReleaseVersionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
