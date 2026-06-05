// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPublishStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppPublishStatusRequest
	GetBizId() *string
	SetDeployOrderId(v int64) *GetAppPublishStatusRequest
	GetDeployOrderId() *int64
	SetWebsiteDomain(v string) *GetAppPublishStatusRequest
	GetWebsiteDomain() *string
}

type GetAppPublishStatusRequest struct {
	// example:
	//
	// WS20250731233102000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 4998717
	DeployOrderId *int64 `json:"DeployOrderId,omitempty" xml:"DeployOrderId,omitempty"`
	// example:
	//
	// www.aliyun.com
	WebsiteDomain *string `json:"WebsiteDomain,omitempty" xml:"WebsiteDomain,omitempty"`
}

func (s GetAppPublishStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppPublishStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAppPublishStatusRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppPublishStatusRequest) GetDeployOrderId() *int64 {
	return s.DeployOrderId
}

func (s *GetAppPublishStatusRequest) GetWebsiteDomain() *string {
	return s.WebsiteDomain
}

func (s *GetAppPublishStatusRequest) SetBizId(v string) *GetAppPublishStatusRequest {
	s.BizId = &v
	return s
}

func (s *GetAppPublishStatusRequest) SetDeployOrderId(v int64) *GetAppPublishStatusRequest {
	s.DeployOrderId = &v
	return s
}

func (s *GetAppPublishStatusRequest) SetWebsiteDomain(v string) *GetAppPublishStatusRequest {
	s.WebsiteDomain = &v
	return s
}

func (s *GetAppPublishStatusRequest) Validate() error {
	return dara.Validate(s)
}
