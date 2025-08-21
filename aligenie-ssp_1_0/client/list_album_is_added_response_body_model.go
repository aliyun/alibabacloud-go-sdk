// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlbumIsAddedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListAlbumIsAddedResponseBody
	GetCode() *int32
	SetMessage(v string) *ListAlbumIsAddedResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAlbumIsAddedResponseBody
	GetRequestId() *string
	SetResult(v []*ListAlbumIsAddedResponseBodyResult) *ListAlbumIsAddedResponseBody
	GetResult() []*ListAlbumIsAddedResponseBodyResult
}

type ListAlbumIsAddedResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BCC85E69-5DA6-197E-A8C1-8A1B19CF781B
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListAlbumIsAddedResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListAlbumIsAddedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumIsAddedResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListAlbumIsAddedResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAlbumIsAddedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlbumIsAddedResponseBody) GetResult() []*ListAlbumIsAddedResponseBodyResult {
	return s.Result
}

func (s *ListAlbumIsAddedResponseBody) SetCode(v int32) *ListAlbumIsAddedResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlbumIsAddedResponseBody) SetMessage(v string) *ListAlbumIsAddedResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlbumIsAddedResponseBody) SetRequestId(v string) *ListAlbumIsAddedResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlbumIsAddedResponseBody) SetResult(v []*ListAlbumIsAddedResponseBodyResult) *ListAlbumIsAddedResponseBody {
	s.Result = v
	return s
}

func (s *ListAlbumIsAddedResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAlbumIsAddedResponseBodyResult struct {
	// example:
	//
	// 51999575
	AlbumId *string `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	// example:
	//
	// false
	IsAdded *string `json:"IsAdded,omitempty" xml:"IsAdded,omitempty"`
}

func (s ListAlbumIsAddedResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumIsAddedResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedResponseBodyResult) GetAlbumId() *string {
	return s.AlbumId
}

func (s *ListAlbumIsAddedResponseBodyResult) GetIsAdded() *string {
	return s.IsAdded
}

func (s *ListAlbumIsAddedResponseBodyResult) SetAlbumId(v string) *ListAlbumIsAddedResponseBodyResult {
	s.AlbumId = &v
	return s
}

func (s *ListAlbumIsAddedResponseBodyResult) SetIsAdded(v string) *ListAlbumIsAddedResponseBodyResult {
	s.IsAdded = &v
	return s
}

func (s *ListAlbumIsAddedResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
