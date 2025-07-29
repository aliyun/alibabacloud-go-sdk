// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKubernetesTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKubernetesTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKubernetesTriggerResponse
	GetStatusCode() *int32
	SetBody(v []*GetKubernetesTriggerResponseBody) *GetKubernetesTriggerResponse
	GetBody() []*GetKubernetesTriggerResponseBody
}

type GetKubernetesTriggerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*GetKubernetesTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s GetKubernetesTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKubernetesTriggerResponse) GoString() string {
	return s.String()
}

func (s *GetKubernetesTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKubernetesTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKubernetesTriggerResponse) GetBody() []*GetKubernetesTriggerResponseBody {
	return s.Body
}

func (s *GetKubernetesTriggerResponse) SetHeaders(v map[string]*string) *GetKubernetesTriggerResponse {
	s.Headers = v
	return s
}

func (s *GetKubernetesTriggerResponse) SetStatusCode(v int32) *GetKubernetesTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKubernetesTriggerResponse) SetBody(v []*GetKubernetesTriggerResponseBody) *GetKubernetesTriggerResponse {
	s.Body = v
	return s
}

func (s *GetKubernetesTriggerResponse) Validate() error {
	return dara.Validate(s)
}

type GetKubernetesTriggerResponseBody struct {
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
	// c259f563386444ebb8d7**
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
	// Token
	//
	// example:
	//
	// eyJhbGci***
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetKubernetesTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKubernetesTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *GetKubernetesTriggerResponseBody) GetId() *string {
	return s.Id
}

func (s *GetKubernetesTriggerResponseBody) GetName() *string {
	return s.Name
}

func (s *GetKubernetesTriggerResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetKubernetesTriggerResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetKubernetesTriggerResponseBody) GetType() *string {
	return s.Type
}

func (s *GetKubernetesTriggerResponseBody) GetAction() *string {
	return s.Action
}

func (s *GetKubernetesTriggerResponseBody) GetToken() *string {
	return s.Token
}

func (s *GetKubernetesTriggerResponseBody) SetId(v string) *GetKubernetesTriggerResponseBody {
	s.Id = &v
	return s
}

func (s *GetKubernetesTriggerResponseBody) SetName(v string) *GetKubernetesTriggerResponseBody {
	s.Name = &v
	return s
}

func (s *GetKubernetesTriggerResponseBody) SetClusterId(v string) *GetKubernetesTriggerResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetKubernetesTriggerResponseBody) SetProjectId(v string) *GetKubernetesTriggerResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetKubernetesTriggerResponseBody) SetType(v string) *GetKubernetesTriggerResponseBody {
	s.Type = &v
	return s
}

func (s *GetKubernetesTriggerResponseBody) SetAction(v string) *GetKubernetesTriggerResponseBody {
	s.Action = &v
	return s
}

func (s *GetKubernetesTriggerResponseBody) SetToken(v string) *GetKubernetesTriggerResponseBody {
	s.Token = &v
	return s
}

func (s *GetKubernetesTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
