// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageResultExtRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInfoType(v string) *DescribeImageResultExtRequest
	GetInfoType() *string
	SetReqId(v string) *DescribeImageResultExtRequest
	GetReqId() *string
}

type DescribeImageResultExtRequest struct {
	// The type of information to obtain. Multiple values are separated by commas. Valid values:
	//
	// - customImage: custom image library hit information
	//
	// - textInImage: text information in the image
	//
	// example:
	//
	// customImage,textInImage
	InfoType *string `json:"InfoType,omitempty" xml:"InfoType,omitempty"`
	// The requestId field returned by the enhanced image moderation API
	//
	// example:
	//
	// 638EDDC65C82AB39319A9F60
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
}

func (s DescribeImageResultExtRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageResultExtRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtRequest) GetInfoType() *string {
	return s.InfoType
}

func (s *DescribeImageResultExtRequest) GetReqId() *string {
	return s.ReqId
}

func (s *DescribeImageResultExtRequest) SetInfoType(v string) *DescribeImageResultExtRequest {
	s.InfoType = &v
	return s
}

func (s *DescribeImageResultExtRequest) SetReqId(v string) *DescribeImageResultExtRequest {
	s.ReqId = &v
	return s
}

func (s *DescribeImageResultExtRequest) Validate() error {
	return dara.Validate(s)
}
