// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultimodalLabelStudioServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteMultimodalLabelStudioServiceRequest
	GetDBClusterId() *string
	SetSourceRegionId(v string) *DeleteMultimodalLabelStudioServiceRequest
	GetSourceRegionId() *string
}

type DeleteMultimodalLabelStudioServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// cn-beijing
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
}

func (s DeleteMultimodalLabelStudioServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultimodalLabelStudioServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultimodalLabelStudioServiceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteMultimodalLabelStudioServiceRequest) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *DeleteMultimodalLabelStudioServiceRequest) SetDBClusterId(v string) *DeleteMultimodalLabelStudioServiceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteMultimodalLabelStudioServiceRequest) SetSourceRegionId(v string) *DeleteMultimodalLabelStudioServiceRequest {
	s.SourceRegionId = &v
	return s
}

func (s *DeleteMultimodalLabelStudioServiceRequest) Validate() error {
	return dara.Validate(s)
}
