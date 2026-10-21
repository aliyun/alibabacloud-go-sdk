// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteImageCacheResponseBody
	GetCode() *string
	SetData(v *DeleteImageCacheResponseBodyData) *DeleteImageCacheResponseBody
	GetData() *DeleteImageCacheResponseBodyData
	SetHttpStatusCode(v int32) *DeleteImageCacheResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteImageCacheResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteImageCacheResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteImageCacheResponseBody
	GetSuccess() *bool
}

type DeleteImageCacheResponseBody struct {
	// The business response code.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result of deleting the image cache.
	//
	// This parameter is required.
	Data *DeleteImageCacheResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is used for troubleshooting.
	//
	// This parameter is required.
	//
	// example:
	//
	// req-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful. This does not mean that the cache creation is complete.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteImageCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteImageCacheResponseBody) GetData() *DeleteImageCacheResponseBodyData {
	return s.Data
}

func (s *DeleteImageCacheResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteImageCacheResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteImageCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteImageCacheResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteImageCacheResponseBody) SetCode(v string) *DeleteImageCacheResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteImageCacheResponseBody) SetData(v *DeleteImageCacheResponseBodyData) *DeleteImageCacheResponseBody {
	s.Data = v
	return s
}

func (s *DeleteImageCacheResponseBody) SetHttpStatusCode(v int32) *DeleteImageCacheResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteImageCacheResponseBody) SetMessage(v string) *DeleteImageCacheResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteImageCacheResponseBody) SetRequestId(v string) *DeleteImageCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageCacheResponseBody) SetSuccess(v bool) *DeleteImageCacheResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteImageCacheResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteImageCacheResponseBodyData struct {
	// Indicates whether the deletion is still in progress. A value of true indicates that the deletion is accepted but not yet complete, and a Get request returns DELETING until ImageCacheNotFound is explicitly returned. A value of false indicates that the current workspace is deleted or no longer exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DeletionPending *bool `json:"deletionPending,omitempty" xml:"deletionPending,omitempty"`
	// The image cache ID, which matches the request. The requested ID is still returned even if the current workspace no longer exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// img-example
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
}

func (s DeleteImageCacheResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageCacheResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheResponseBodyData) GetDeletionPending() *bool {
	return s.DeletionPending
}

func (s *DeleteImageCacheResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *DeleteImageCacheResponseBodyData) SetDeletionPending(v bool) *DeleteImageCacheResponseBodyData {
	s.DeletionPending = &v
	return s
}

func (s *DeleteImageCacheResponseBodyData) SetImageId(v string) *DeleteImageCacheResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *DeleteImageCacheResponseBodyData) Validate() error {
	return dara.Validate(s)
}
