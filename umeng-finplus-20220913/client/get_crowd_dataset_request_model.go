// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrowdDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *GetCrowdDatasetRequestBody) *GetCrowdDatasetRequest
	GetBody() *GetCrowdDatasetRequestBody
}

type GetCrowdDatasetRequest struct {
	Body *GetCrowdDatasetRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s GetCrowdDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCrowdDatasetRequest) GoString() string {
	return s.String()
}

func (s *GetCrowdDatasetRequest) GetBody() *GetCrowdDatasetRequestBody {
	return s.Body
}

func (s *GetCrowdDatasetRequest) SetBody(v *GetCrowdDatasetRequestBody) *GetCrowdDatasetRequest {
	s.Body = v
	return s
}

func (s *GetCrowdDatasetRequest) Validate() error {
	return dara.Validate(s)
}

type GetCrowdDatasetRequestBody struct {
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
}

func (s GetCrowdDatasetRequestBody) String() string {
	return dara.Prettify(s)
}

func (s GetCrowdDatasetRequestBody) GoString() string {
	return s.String()
}

func (s *GetCrowdDatasetRequestBody) GetAppId() *string {
	return s.AppId
}

func (s *GetCrowdDatasetRequestBody) SetAppId(v string) *GetCrowdDatasetRequestBody {
	s.AppId = &v
	return s
}

func (s *GetCrowdDatasetRequestBody) Validate() error {
	return dara.Validate(s)
}
