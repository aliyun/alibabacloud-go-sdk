// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserTagMetaListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryUserTagMetaListResponseBody
	GetRequestId() *string
	SetResult(v []*QueryUserTagMetaListResponseBodyResult) *QueryUserTagMetaListResponseBody
	GetResult() []*QueryUserTagMetaListResponseBodyResult
	SetSuccess(v bool) *QueryUserTagMetaListResponseBody
	GetSuccess() *bool
}

type QueryUserTagMetaListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns a list of user tags in an organization.
	Result []*QueryUserTagMetaListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserTagMetaListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserTagMetaListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserTagMetaListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserTagMetaListResponseBody) GetResult() []*QueryUserTagMetaListResponseBodyResult {
	return s.Result
}

func (s *QueryUserTagMetaListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUserTagMetaListResponseBody) SetRequestId(v string) *QueryUserTagMetaListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserTagMetaListResponseBody) SetResult(v []*QueryUserTagMetaListResponseBodyResult) *QueryUserTagMetaListResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserTagMetaListResponseBody) SetSuccess(v bool) *QueryUserTagMetaListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserTagMetaListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryUserTagMetaListResponseBodyResult struct {
	// The description of the tag.
	//
	// example:
	//
	// Used to distinguish some positions
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// The ID of the label.
	//
	// example:
	//
	// pop_001
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The name of the tag.
	//
	// example:
	//
	// Position
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s QueryUserTagMetaListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryUserTagMetaListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserTagMetaListResponseBodyResult) GetTagDescription() *string {
	return s.TagDescription
}

func (s *QueryUserTagMetaListResponseBodyResult) GetTagId() *string {
	return s.TagId
}

func (s *QueryUserTagMetaListResponseBodyResult) GetTagName() *string {
	return s.TagName
}

func (s *QueryUserTagMetaListResponseBodyResult) SetTagDescription(v string) *QueryUserTagMetaListResponseBodyResult {
	s.TagDescription = &v
	return s
}

func (s *QueryUserTagMetaListResponseBodyResult) SetTagId(v string) *QueryUserTagMetaListResponseBodyResult {
	s.TagId = &v
	return s
}

func (s *QueryUserTagMetaListResponseBodyResult) SetTagName(v string) *QueryUserTagMetaListResponseBodyResult {
	s.TagName = &v
	return s
}

func (s *QueryUserTagMetaListResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
