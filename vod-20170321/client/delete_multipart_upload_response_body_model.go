// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultipartUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMultipartUploadResponseBody
	GetRequestId() *string
}

type DeleteMultipartUploadResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMultipartUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultipartUploadResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultipartUploadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMultipartUploadResponseBody) SetRequestId(v string) *DeleteMultipartUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMultipartUploadResponseBody) Validate() error {
	return dara.Validate(s)
}
