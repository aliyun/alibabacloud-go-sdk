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
	return dara.Validate(s)
}

type DescribeTriggerResponseBody struct {
	// The ID of the trigger.
	//
	// example:
	//
	// 1234
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the trigger.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the associated cluster.
	//
	// example:
	//
	// c259f563386444ebb8d7****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The name of the project.
	//
	// The name consists of the namespace where the application is deployed and the name of the application. The format is `${namespace}/${name}`. Example: default/test-app.
	//
	// example:
	//
	// default/test-app
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// The type of trigger.
	//
	// Valid values:
	//
	// 	- `deployment`: performs actions on Deployments.
	//
	// 	- `application`: performs actions on applications that are deployed in Application Center.
	//
	// Default value: `deployment`.
	//
	// example:
	//
	// deployment
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The action that the trigger performs. The value is set to redeploy.
	//
	// `redeploy`: redeploys the resource specified by project_id.
	//
	// example:
	//
	// redeploy
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The token information.
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
