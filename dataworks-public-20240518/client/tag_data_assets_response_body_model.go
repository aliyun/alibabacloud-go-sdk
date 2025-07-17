// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagDataAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TagDataAssetsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *TagDataAssetsResponseBody
	GetSuccess() *string
}

type TagDataAssetsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TagDataAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagDataAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *TagDataAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagDataAssetsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *TagDataAssetsResponseBody) SetRequestId(v string) *TagDataAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagDataAssetsResponseBody) SetSuccess(v string) *TagDataAssetsResponseBody {
	s.Success = &v
	return s
}

func (s *TagDataAssetsResponseBody) Validate() error {
	return dara.Validate(s)
}
