// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserTagValueListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryUserTagValueListResponseBody
	GetRequestId() *string
	SetResult(v []*QueryUserTagValueListResponseBodyResult) *QueryUserTagValueListResponseBody
	GetResult() []*QueryUserTagValueListResponseBodyResult
	SetSuccess(v bool) *QueryUserTagValueListResponseBody
	GetSuccess() *bool
}

type QueryUserTagValueListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request for a list of user tags and their values.
	Result []*QueryUserTagValueListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserTagValueListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserTagValueListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserTagValueListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserTagValueListResponseBody) GetResult() []*QueryUserTagValueListResponseBodyResult {
	return s.Result
}

func (s *QueryUserTagValueListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUserTagValueListResponseBody) SetRequestId(v string) *QueryUserTagValueListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserTagValueListResponseBody) SetResult(v []*QueryUserTagValueListResponseBodyResult) *QueryUserTagValueListResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserTagValueListResponseBody) SetSuccess(v bool) *QueryUserTagValueListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserTagValueListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryUserTagValueListResponseBodyResult struct {
	// Tag ID.
	//
	// example:
	//
	// pop_001
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// Tag name.
	//
	// example:
	//
	// Position
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// Supervisor
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s QueryUserTagValueListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryUserTagValueListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserTagValueListResponseBodyResult) GetTagId() *string {
	return s.TagId
}

func (s *QueryUserTagValueListResponseBodyResult) GetTagName() *string {
	return s.TagName
}

func (s *QueryUserTagValueListResponseBodyResult) GetTagValue() *string {
	return s.TagValue
}

func (s *QueryUserTagValueListResponseBodyResult) SetTagId(v string) *QueryUserTagValueListResponseBodyResult {
	s.TagId = &v
	return s
}

func (s *QueryUserTagValueListResponseBodyResult) SetTagName(v string) *QueryUserTagValueListResponseBodyResult {
	s.TagName = &v
	return s
}

func (s *QueryUserTagValueListResponseBodyResult) SetTagValue(v string) *QueryUserTagValueListResponseBodyResult {
	s.TagValue = &v
	return s
}

func (s *QueryUserTagValueListResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
