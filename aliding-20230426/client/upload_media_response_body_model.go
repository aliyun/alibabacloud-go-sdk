// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *UploadMediaResponseBody
	GetMediaId() *string
	SetRequestId(v string) *UploadMediaResponseBody
	GetRequestId() *string
}

type UploadMediaResponseBody struct {
	// example:
	//
	// abcd1234
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// requestId
	//
	// example:
	//
	// abcd-1234-wxyz-9876
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UploadMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadMediaResponseBody) GoString() string {
	return s.String()
}

func (s *UploadMediaResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *UploadMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadMediaResponseBody) SetMediaId(v string) *UploadMediaResponseBody {
	s.MediaId = &v
	return s
}

func (s *UploadMediaResponseBody) SetRequestId(v string) *UploadMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadMediaResponseBody) Validate() error {
	return dara.Validate(s)
}
