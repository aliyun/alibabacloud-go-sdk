// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLikeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendLikeResponseBody
	GetRequestId() *string
	SetResult(v *SendLikeResponseBodyResult) *SendLikeResponseBody
	GetResult() *SendLikeResponseBodyResult
}

type SendLikeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *SendLikeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s SendLikeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendLikeResponseBody) GoString() string {
	return s.String()
}

func (s *SendLikeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendLikeResponseBody) GetResult() *SendLikeResponseBodyResult {
	return s.Result
}

func (s *SendLikeResponseBody) SetRequestId(v string) *SendLikeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendLikeResponseBody) SetResult(v *SendLikeResponseBodyResult) *SendLikeResponseBody {
	s.Result = v
	return s
}

func (s *SendLikeResponseBody) Validate() error {
	return dara.Validate(s)
}

type SendLikeResponseBodyResult struct {
	// The number of likes.
	//
	// example:
	//
	// 10
	LikeCount *int32 `json:"LikeCount,omitempty" xml:"LikeCount,omitempty"`
}

func (s SendLikeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s SendLikeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendLikeResponseBodyResult) GetLikeCount() *int32 {
	return s.LikeCount
}

func (s *SendLikeResponseBodyResult) SetLikeCount(v int32) *SendLikeResponseBodyResult {
	s.LikeCount = &v
	return s
}

func (s *SendLikeResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
