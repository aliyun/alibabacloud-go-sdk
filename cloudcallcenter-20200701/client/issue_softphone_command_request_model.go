// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIssueSoftphoneCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *IssueSoftphoneCommandRequest
	GetData() *string
}

type IssueSoftphoneCommandRequest struct {
	// This parameter is required.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s IssueSoftphoneCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s IssueSoftphoneCommandRequest) GoString() string {
	return s.String()
}

func (s *IssueSoftphoneCommandRequest) GetData() *string {
	return s.Data
}

func (s *IssueSoftphoneCommandRequest) SetData(v string) *IssueSoftphoneCommandRequest {
	s.Data = &v
	return s
}

func (s *IssueSoftphoneCommandRequest) Validate() error {
	return dara.Validate(s)
}
