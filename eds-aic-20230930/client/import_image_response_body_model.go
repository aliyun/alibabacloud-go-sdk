// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ImportImageResponseBody
	GetImageId() *string
	SetRequestId(v string) *ImportImageResponseBody
	GetRequestId() *string
}

type ImportImageResponseBody struct {
	// ImageId
	//
	// example:
	//
	// imgc-06zyt9m93zwax****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportImageResponseBody) GoString() string {
	return s.String()
}

func (s *ImportImageResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *ImportImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportImageResponseBody) SetImageId(v string) *ImportImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *ImportImageResponseBody) SetRequestId(v string) *ImportImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportImageResponseBody) Validate() error {
	return dara.Validate(s)
}
