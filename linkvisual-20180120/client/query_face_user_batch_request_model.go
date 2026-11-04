// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsolationId(v string) *QueryFaceUserBatchRequest
	GetIsolationId() *string
	SetUserIdList(v []*string) *QueryFaceUserBatchRequest
	GetUserIdList() []*string
}

type QueryFaceUserBatchRequest struct {
	// The business isolation ID, which is used to isolate data between different business applications under the same Alibaba Cloud account. Keep this value consistent with the **IsolationId*	- set when calling the [AddFaceUser](https://help.aliyun.com/document_detail/288338.html) operation to add users.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZheJiangHZ
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The list of user IDs. Pass the user IDs that were obtained and saved locally after calling the [AddFaceUser](https://help.aliyun.com/document_detail/288338.html) operation to add users.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["xqymlby****4whye","6qyyb69zqara****","64nbfx8****hzmy4"]
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s QueryFaceUserBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserBatchRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserBatchRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *QueryFaceUserBatchRequest) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *QueryFaceUserBatchRequest) SetIsolationId(v string) *QueryFaceUserBatchRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceUserBatchRequest) SetUserIdList(v []*string) *QueryFaceUserBatchRequest {
	s.UserIdList = v
	return s
}

func (s *QueryFaceUserBatchRequest) Validate() error {
	return dara.Validate(s)
}
