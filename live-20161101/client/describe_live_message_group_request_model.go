// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveMessageGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeLiveMessageGroupRequest
	GetAppId() *string
	SetDataCenter(v string) *DescribeLiveMessageGroupRequest
	GetDataCenter() *string
	SetGroupId(v string) *DescribeLiveMessageGroupRequest
	GetGroupId() *string
}

type DescribeLiveMessageGroupRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// coims-pre
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The ID of the interactive messaging group whose information you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeLiveMessageGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveMessageGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveMessageGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeLiveMessageGroupRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *DescribeLiveMessageGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeLiveMessageGroupRequest) SetAppId(v string) *DescribeLiveMessageGroupRequest {
	s.AppId = &v
	return s
}

func (s *DescribeLiveMessageGroupRequest) SetDataCenter(v string) *DescribeLiveMessageGroupRequest {
	s.DataCenter = &v
	return s
}

func (s *DescribeLiveMessageGroupRequest) SetGroupId(v string) *DescribeLiveMessageGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeLiveMessageGroupRequest) Validate() error {
	return dara.Validate(s)
}
