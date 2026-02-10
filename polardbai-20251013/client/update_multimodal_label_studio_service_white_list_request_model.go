// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultimodalLabelStudioServiceWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UpdateMultimodalLabelStudioServiceWhiteListRequest
	GetDBClusterId() *string
	SetWhiteList(v []*string) *UpdateMultimodalLabelStudioServiceWhiteListRequest
	GetWhiteList() []*string
}

type UpdateMultimodalLabelStudioServiceWhiteListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string   `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	WhiteList   []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s UpdateMultimodalLabelStudioServiceWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultimodalLabelStudioServiceWhiteListRequest) GoString() string {
	return s.String()
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListRequest) GetWhiteList() []*string {
	return s.WhiteList
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListRequest) SetDBClusterId(v string) *UpdateMultimodalLabelStudioServiceWhiteListRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListRequest) SetWhiteList(v []*string) *UpdateMultimodalLabelStudioServiceWhiteListRequest {
	s.WhiteList = v
	return s
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
