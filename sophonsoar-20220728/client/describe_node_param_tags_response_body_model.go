// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeParamTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParamReferredPaths(v []*DescribeNodeParamTagsResponseBodyParamReferredPaths) *DescribeNodeParamTagsResponseBody
	GetParamReferredPaths() []*DescribeNodeParamTagsResponseBodyParamReferredPaths
	SetRequestId(v string) *DescribeNodeParamTagsResponseBody
	GetRequestId() *string
}

type DescribeNodeParamTagsResponseBody struct {
	// The configuration of the recommended path.
	ParamReferredPaths []*DescribeNodeParamTagsResponseBodyParamReferredPaths `json:"ParamReferredPaths,omitempty" xml:"ParamReferredPaths,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6BE94351-712A-505D-A40A-BC77CC8254A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNodeParamTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeParamTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeParamTagsResponseBody) GetParamReferredPaths() []*DescribeNodeParamTagsResponseBodyParamReferredPaths {
	return s.ParamReferredPaths
}

func (s *DescribeNodeParamTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNodeParamTagsResponseBody) SetParamReferredPaths(v []*DescribeNodeParamTagsResponseBodyParamReferredPaths) *DescribeNodeParamTagsResponseBody {
	s.ParamReferredPaths = v
	return s
}

func (s *DescribeNodeParamTagsResponseBody) SetRequestId(v string) *DescribeNodeParamTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeParamTagsResponseBody) Validate() error {
	if s.ParamReferredPaths != nil {
		for _, item := range s.ParamReferredPaths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNodeParamTagsResponseBodyParamReferredPaths struct {
	// The name of the upstream node.
	//
	// example:
	//
	// DataFormat_1
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// The paths.
	ReferredPath []*string `json:"ReferredPath,omitempty" xml:"ReferredPath,omitempty" type:"Repeated"`
}

func (s DescribeNodeParamTagsResponseBodyParamReferredPaths) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeParamTagsResponseBodyParamReferredPaths) GoString() string {
	return s.String()
}

func (s *DescribeNodeParamTagsResponseBodyParamReferredPaths) GetParamName() *string {
	return s.ParamName
}

func (s *DescribeNodeParamTagsResponseBodyParamReferredPaths) GetReferredPath() []*string {
	return s.ReferredPath
}

func (s *DescribeNodeParamTagsResponseBodyParamReferredPaths) SetParamName(v string) *DescribeNodeParamTagsResponseBodyParamReferredPaths {
	s.ParamName = &v
	return s
}

func (s *DescribeNodeParamTagsResponseBodyParamReferredPaths) SetReferredPath(v []*string) *DescribeNodeParamTagsResponseBodyParamReferredPaths {
	s.ReferredPath = v
	return s
}

func (s *DescribeNodeParamTagsResponseBodyParamReferredPaths) Validate() error {
	return dara.Validate(s)
}
