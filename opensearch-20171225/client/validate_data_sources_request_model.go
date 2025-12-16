// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DataSource) *ValidateDataSourcesRequest
	GetBody() *DataSource
}

type ValidateDataSourcesRequest struct {
	// The request parameter. For more information, see [DataSource](https://help.aliyun.com/document_detail/170005.html).
	Body *DataSource `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ValidateDataSourcesRequest) GetBody() *DataSource {
	return s.Body
}

func (s *ValidateDataSourcesRequest) SetBody(v *DataSource) *ValidateDataSourcesRequest {
	s.Body = v
	return s
}

func (s *ValidateDataSourcesRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
