// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeTranscodeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetEdgeTranscodeTemplateRequest
	GetClusterId() *string
	SetOwnerId(v int64) *GetEdgeTranscodeTemplateRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetEdgeTranscodeTemplateRequest
	GetRegionId() *string
	SetTemplateId(v string) *GetEdgeTranscodeTemplateRequest
	GetTemplateId() *string
}

type GetEdgeTranscodeTemplateRequest struct {
	// The ID of the data center.
	//
	// This parameter is required.
	//
	// example:
	//
	// ******3b-4d18-395c-8106-ff21a6******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287666****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetEdgeTranscodeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeTranscodeTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeTranscodeTemplateRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetEdgeTranscodeTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetEdgeTranscodeTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEdgeTranscodeTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetEdgeTranscodeTemplateRequest) SetClusterId(v string) *GetEdgeTranscodeTemplateRequest {
	s.ClusterId = &v
	return s
}

func (s *GetEdgeTranscodeTemplateRequest) SetOwnerId(v int64) *GetEdgeTranscodeTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *GetEdgeTranscodeTemplateRequest) SetRegionId(v string) *GetEdgeTranscodeTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetEdgeTranscodeTemplateRequest) SetTemplateId(v string) *GetEdgeTranscodeTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetEdgeTranscodeTemplateRequest) Validate() error {
	return dara.Validate(s)
}
