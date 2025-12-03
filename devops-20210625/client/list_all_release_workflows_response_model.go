// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllReleaseWorkflowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllReleaseWorkflowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllReleaseWorkflowsResponse
	GetStatusCode() *int32
	SetBody(v []*ListAllReleaseWorkflowsResponseBody) *ListAllReleaseWorkflowsResponse
	GetBody() []*ListAllReleaseWorkflowsResponseBody
}

type ListAllReleaseWorkflowsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*ListAllReleaseWorkflowsResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListAllReleaseWorkflowsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllReleaseWorkflowsResponse) GoString() string {
	return s.String()
}

func (s *ListAllReleaseWorkflowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllReleaseWorkflowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllReleaseWorkflowsResponse) GetBody() []*ListAllReleaseWorkflowsResponseBody {
	return s.Body
}

func (s *ListAllReleaseWorkflowsResponse) SetHeaders(v map[string]*string) *ListAllReleaseWorkflowsResponse {
	s.Headers = v
	return s
}

func (s *ListAllReleaseWorkflowsResponse) SetStatusCode(v int32) *ListAllReleaseWorkflowsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllReleaseWorkflowsResponse) SetBody(v []*ListAllReleaseWorkflowsResponseBody) *ListAllReleaseWorkflowsResponse {
	s.Body = v
	return s
}

func (s *ListAllReleaseWorkflowsResponse) Validate() error {
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

type ListAllReleaseWorkflowsResponseBody struct {
	// example:
	//
	// testApp
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// example:
	//
	// ce51b31b996246ecaf874736838360b2
	Sn   *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	Order         *string                                             `json:"order,omitempty" xml:"order,omitempty"`
	ReleaseStages []*ListAllReleaseWorkflowsResponseBodyReleaseStages `json:"releaseStages,omitempty" xml:"releaseStages,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
}

func (s ListAllReleaseWorkflowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllReleaseWorkflowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllReleaseWorkflowsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAllReleaseWorkflowsResponseBody) GetSn() *string {
	return s.Sn
}

func (s *ListAllReleaseWorkflowsResponseBody) GetName() *string {
	return s.Name
}

func (s *ListAllReleaseWorkflowsResponseBody) GetOrder() *string {
	return s.Order
}

func (s *ListAllReleaseWorkflowsResponseBody) GetReleaseStages() []*ListAllReleaseWorkflowsResponseBodyReleaseStages {
	return s.ReleaseStages
}

func (s *ListAllReleaseWorkflowsResponseBody) GetNote() *string {
	return s.Note
}

func (s *ListAllReleaseWorkflowsResponseBody) SetAppName(v string) *ListAllReleaseWorkflowsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBody) SetSn(v string) *ListAllReleaseWorkflowsResponseBody {
	s.Sn = &v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBody) SetName(v string) *ListAllReleaseWorkflowsResponseBody {
	s.Name = &v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBody) SetOrder(v string) *ListAllReleaseWorkflowsResponseBody {
	s.Order = &v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBody) SetReleaseStages(v []*ListAllReleaseWorkflowsResponseBodyReleaseStages) *ListAllReleaseWorkflowsResponseBody {
	s.ReleaseStages = v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBody) SetNote(v string) *ListAllReleaseWorkflowsResponseBody {
	s.Note = &v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBody) Validate() error {
	if s.ReleaseStages != nil {
		for _, item := range s.ReleaseStages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAllReleaseWorkflowsResponseBodyReleaseStages struct {
	// example:
	//
	// testApp
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 5aa8cc67e75e41bf9dddeb708775bcc3
	Sn *string `json:"sn,omitempty" xml:"sn,omitempty"`
	// example:
	//
	// ce51b31b996246ecaf874736838360b2
	ReleaseWorkflowSn *string `json:"releaseWorkflowSn,omitempty" xml:"releaseWorkflowSn,omitempty"`
	// example:
	//
	// 1
	Order          *string                                                           `json:"order,omitempty" xml:"order,omitempty"`
	VariableGroups []*ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups `json:"variableGroups,omitempty" xml:"variableGroups,omitempty" type:"Repeated"`
}

func (s ListAllReleaseWorkflowsResponseBodyReleaseStages) String() string {
	return dara.Prettify(s)
}

func (s ListAllReleaseWorkflowsResponseBodyReleaseStages) GoString() string {
	return s.String()
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStages) GetAppName() *string {
	return s.AppName
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStages) GetName() *string {
	return s.Name
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStages) GetSn() *string {
	return s.Sn
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStages) GetReleaseWorkflowSn() *string {
	return s.ReleaseWorkflowSn
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStages) GetOrder() *string {
	return s.Order
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStages) GetVariableGroups() []*ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups {
	return s.VariableGroups
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStages) SetAppName(v string) *ListAllReleaseWorkflowsResponseBodyReleaseStages {
	s.AppName = &v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStages) SetName(v string) *ListAllReleaseWorkflowsResponseBodyReleaseStages {
	s.Name = &v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStages) SetSn(v string) *ListAllReleaseWorkflowsResponseBodyReleaseStages {
	s.Sn = &v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStages) SetReleaseWorkflowSn(v string) *ListAllReleaseWorkflowsResponseBodyReleaseStages {
	s.ReleaseWorkflowSn = &v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStages) SetOrder(v string) *ListAllReleaseWorkflowsResponseBodyReleaseStages {
	s.Order = &v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStages) SetVariableGroups(v []*ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups) *ListAllReleaseWorkflowsResponseBodyReleaseStages {
	s.VariableGroups = v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStages) Validate() error {
	if s.VariableGroups != nil {
		for _, item := range s.VariableGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups struct {
	// example:
	//
	// dev
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// APP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups) String() string {
	return dara.Prettify(s)
}

func (s ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups) GoString() string {
	return s.String()
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups) GetName() *string {
	return s.Name
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups) GetType() *string {
	return s.Type
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups) SetName(v string) *ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups {
	s.Name = &v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups) SetDisplayName(v string) *ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups {
	s.DisplayName = &v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups) SetType(v string) *ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups {
	s.Type = &v
	return s
}

func (s *ListAllReleaseWorkflowsResponseBodyReleaseStagesVariableGroups) Validate() error {
	return dara.Validate(s)
}
