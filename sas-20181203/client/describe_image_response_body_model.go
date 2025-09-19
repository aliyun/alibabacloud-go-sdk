// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeImageResponseBodyData) *DescribeImageResponseBody
	GetData() *DescribeImageResponseBodyData
	SetRequestId(v string) *DescribeImageResponseBody
	GetRequestId() *string
}

type DescribeImageResponseBody struct {
	// The information about the image digest.
	Data *DescribeImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageResponseBody) GetData() *DescribeImageResponseBodyData {
	return s.Data
}

func (s *DescribeImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageResponseBody) SetData(v *DescribeImageResponseBodyData) *DescribeImageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageResponseBody) SetRequestId(v string) *DescribeImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageResponseBodyData struct {
	// The digest value of the image.
	//
	// example:
	//
	// 0afb98d97f1a4030782fcf47e186909e5ad957bcc182d8be70334e0684b2****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
}

func (s DescribeImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageResponseBodyData) GetDigest() *string {
	return s.Digest
}

func (s *DescribeImageResponseBodyData) SetDigest(v string) *DescribeImageResponseBodyData {
	s.Digest = &v
	return s
}

func (s *DescribeImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
