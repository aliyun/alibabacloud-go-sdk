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
	// The content of the information to be obtained. Multiple values are separated by commas.
	//
	// example:
	//
	// customImage,textInImage
	InfoType *string `json:"InfoType,omitempty" xml:"InfoType,omitempty"`
	// The reqId field returned by the Url Async Moderation API.
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
