// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveMessageAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeLiveMessageAppRequest
	GetAppId() *string
	SetDataCenter(v string) *DescribeLiveMessageAppRequest
	GetDataCenter() *string
}

type DescribeLiveMessageAppRequest struct {
	// The ID of the interactive messaging application to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
}

func (s DescribeLiveMessageAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveMessageAppRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveMessageAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeLiveMessageAppRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *DescribeLiveMessageAppRequest) SetAppId(v string) *DescribeLiveMessageAppRequest {
	s.AppId = &v
	return s
}

func (s *DescribeLiveMessageAppRequest) SetDataCenter(v string) *DescribeLiveMessageAppRequest {
	s.DataCenter = &v
	return s
}

func (s *DescribeLiveMessageAppRequest) Validate() error {
	return dara.Validate(s)
}
