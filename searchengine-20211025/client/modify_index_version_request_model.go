// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIndexVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*ModifyIndexVersionRequestBody) *ModifyIndexVersionRequest
	GetBody() []*ModifyIndexVersionRequestBody
}

type ModifyIndexVersionRequest struct {
	// The request body.
	//
	// example:
	//
	// {}
	Body []*ModifyIndexVersionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ModifyIndexVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexVersionRequest) GoString() string {
	return s.String()
}

func (s *ModifyIndexVersionRequest) GetBody() []*ModifyIndexVersionRequestBody {
	return s.Body
}

func (s *ModifyIndexVersionRequest) SetBody(v []*ModifyIndexVersionRequestBody) *ModifyIndexVersionRequest {
	s.Body = v
	return s
}

func (s *ModifyIndexVersionRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyIndexVersionRequestBody struct {
	// The deployment ID of the data source.
	//
	// example:
	//
	// 277
	BuildDeployId *string `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	// The index name.
	//
	// example:
	//
	// main_index
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// The index version.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ModifyIndexVersionRequestBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexVersionRequestBody) GoString() string {
	return s.String()
}

func (s *ModifyIndexVersionRequestBody) GetBuildDeployId() *string {
	return s.BuildDeployId
}

func (s *ModifyIndexVersionRequestBody) GetIndexName() *string {
	return s.IndexName
}

func (s *ModifyIndexVersionRequestBody) GetVersion() *string {
	return s.Version
}

func (s *ModifyIndexVersionRequestBody) SetBuildDeployId(v string) *ModifyIndexVersionRequestBody {
	s.BuildDeployId = &v
	return s
}

func (s *ModifyIndexVersionRequestBody) SetIndexName(v string) *ModifyIndexVersionRequestBody {
	s.IndexName = &v
	return s
}

func (s *ModifyIndexVersionRequestBody) SetVersion(v string) *ModifyIndexVersionRequestBody {
	s.Version = &v
	return s
}

func (s *ModifyIndexVersionRequestBody) Validate() error {
	return dara.Validate(s)
}
