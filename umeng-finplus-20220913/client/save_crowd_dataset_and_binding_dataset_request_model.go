// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveCrowdDatasetAndBindingDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SaveCrowdDatasetAndBindingDatasetRequestBody) *SaveCrowdDatasetAndBindingDatasetRequest
	GetBody() *SaveCrowdDatasetAndBindingDatasetRequestBody
}

type SaveCrowdDatasetAndBindingDatasetRequest struct {
	Body *SaveCrowdDatasetAndBindingDatasetRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s SaveCrowdDatasetAndBindingDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveCrowdDatasetAndBindingDatasetRequest) GoString() string {
	return s.String()
}

func (s *SaveCrowdDatasetAndBindingDatasetRequest) GetBody() *SaveCrowdDatasetAndBindingDatasetRequestBody {
	return s.Body
}

func (s *SaveCrowdDatasetAndBindingDatasetRequest) SetBody(v *SaveCrowdDatasetAndBindingDatasetRequestBody) *SaveCrowdDatasetAndBindingDatasetRequest {
	s.Body = v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetRequest) Validate() error {
	return dara.Validate(s)
}

type SaveCrowdDatasetAndBindingDatasetRequestBody struct {
	// This parameter is required.
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	DatasetIds  []*string `json:"datasetIds,omitempty" xml:"datasetIds,omitempty" type:"Repeated"`
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s SaveCrowdDatasetAndBindingDatasetRequestBody) String() string {
	return dara.Prettify(s)
}

func (s SaveCrowdDatasetAndBindingDatasetRequestBody) GoString() string {
	return s.String()
}

func (s *SaveCrowdDatasetAndBindingDatasetRequestBody) GetAppId() *string {
	return s.AppId
}

func (s *SaveCrowdDatasetAndBindingDatasetRequestBody) GetDatasetIds() []*string {
	return s.DatasetIds
}

func (s *SaveCrowdDatasetAndBindingDatasetRequestBody) GetDescription() *string {
	return s.Description
}

func (s *SaveCrowdDatasetAndBindingDatasetRequestBody) GetName() *string {
	return s.Name
}

func (s *SaveCrowdDatasetAndBindingDatasetRequestBody) SetAppId(v string) *SaveCrowdDatasetAndBindingDatasetRequestBody {
	s.AppId = &v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetRequestBody) SetDatasetIds(v []*string) *SaveCrowdDatasetAndBindingDatasetRequestBody {
	s.DatasetIds = v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetRequestBody) SetDescription(v string) *SaveCrowdDatasetAndBindingDatasetRequestBody {
	s.Description = &v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetRequestBody) SetName(v string) *SaveCrowdDatasetAndBindingDatasetRequestBody {
	s.Name = &v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetRequestBody) Validate() error {
	return dara.Validate(s)
}
