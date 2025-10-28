// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetK8sApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetK8sApplicationRequest
	GetAppId() *string
	SetFrom(v string) *GetK8sApplicationRequest
	GetFrom() *string
}

type GetK8sApplicationRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a166fbd-****-4f98-a286-781659d9****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The source from which data is queried.
	//
	// 	- If you leave this parameter empty, a common query is performed.
	//
	// 	- If you set the value to deploy, you query application information from the deployment page.
	//
	// example:
	//
	// deploy
	From *string `json:"From,omitempty" xml:"From,omitempty"`
}

func (s GetK8sApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetK8sApplicationRequest) GetFrom() *string {
	return s.From
}

func (s *GetK8sApplicationRequest) SetAppId(v string) *GetK8sApplicationRequest {
	s.AppId = &v
	return s
}

func (s *GetK8sApplicationRequest) SetFrom(v string) *GetK8sApplicationRequest {
	s.From = &v
	return s
}

func (s *GetK8sApplicationRequest) Validate() error {
	return dara.Validate(s)
}
