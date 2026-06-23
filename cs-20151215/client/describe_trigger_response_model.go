// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTriggerResponse
	GetStatusCode() *int32
	SetBody(v []*DescribeTriggerResponseBody) *DescribeTriggerResponse
	GetBody() []*DescribeTriggerResponseBody
}

type DescribeTriggerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*DescribeTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s DescribeTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTriggerResponse) GoString() string {
	return s.String()
}

func (s *DescribeTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTriggerResponse) GetBody() []*DescribeTriggerResponseBody {
	return s.Body
}

func (s *DescribeTriggerResponse) SetHeaders(v map[string]*string) *DescribeTriggerResponse {
	s.Headers = v
	return s
}

func (s *DescribeTriggerResponse) SetStatusCode(v int32) *DescribeTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTriggerResponse) SetBody(v []*DescribeTriggerResponseBody) *DescribeTriggerResponse {
	s.Body = v
	return s
}

func (s *DescribeTriggerResponse) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTriggerResponseBody struct {
	// 触发器ID。
	//
	// example:
	//
	// 75973497486******
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 触发器名称。
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 集群ID。
	//
	// example:
	//
	// c259f563386444ebb8d7****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 触发器项目名称。
	//
	// 由应用所在命名空间及应用名称组成，格式为`${namespace}/${name}`，取值示例：default/test-app。
	//
	// example:
	//
	// default/test-app
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 触发器类型。
	//
	// 取值：
	//
	// - `deployment`：针对无状态应用的触发器。
	//
	// - `application`：针对应用中心应用的触发器。
	//
	// example:
	//
	// deployment
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 触发器行为，取值：
	//
	// `redeploy`: 重新部署应用。
	//
	// example:
	//
	// redeploy
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// Token信息。
	//
	// example:
	//
	// eyJhbGci***
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s DescribeTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTriggerResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribeTriggerResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeTriggerResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeTriggerResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeTriggerResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeTriggerResponseBody) GetAction() *string {
	return s.Action
}

func (s *DescribeTriggerResponseBody) GetToken() *string {
	return s.Token
}

func (s *DescribeTriggerResponseBody) SetId(v string) *DescribeTriggerResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeTriggerResponseBody) SetName(v string) *DescribeTriggerResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeTriggerResponseBody) SetClusterId(v string) *DescribeTriggerResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeTriggerResponseBody) SetProjectId(v string) *DescribeTriggerResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeTriggerResponseBody) SetType(v string) *DescribeTriggerResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeTriggerResponseBody) SetAction(v string) *DescribeTriggerResponseBody {
	s.Action = &v
	return s
}

func (s *DescribeTriggerResponseBody) SetToken(v string) *DescribeTriggerResponseBody {
	s.Token = &v
	return s
}

func (s *DescribeTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
