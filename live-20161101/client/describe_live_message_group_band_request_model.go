// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveMessageGroupBandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeLiveMessageGroupBandRequest
	GetAppId() *string
	SetDataCenter(v string) *DescribeLiveMessageGroupBandRequest
	GetDataCenter() *string
	SetGroupId(v string) *DescribeLiveMessageGroupBandRequest
	GetGroupId() *string
}

type DescribeLiveMessageGroupBandRequest struct {
	// The application ID.
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
	// The ID of the interactive messaging group.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeLiveMessageGroupBandRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveMessageGroupBandRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveMessageGroupBandRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeLiveMessageGroupBandRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *DescribeLiveMessageGroupBandRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeLiveMessageGroupBandRequest) SetAppId(v string) *DescribeLiveMessageGroupBandRequest {
	s.AppId = &v
	return s
}

func (s *DescribeLiveMessageGroupBandRequest) SetDataCenter(v string) *DescribeLiveMessageGroupBandRequest {
	s.DataCenter = &v
	return s
}

func (s *DescribeLiveMessageGroupBandRequest) SetGroupId(v string) *DescribeLiveMessageGroupBandRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeLiveMessageGroupBandRequest) Validate() error {
	return dara.Validate(s)
}
