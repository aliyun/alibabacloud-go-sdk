// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelResourceExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelResourceExportTaskRequest
	GetClientToken() *string
}

type CancelResourceExportTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CancelResourceExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelResourceExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelResourceExportTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelResourceExportTaskRequest) SetClientToken(v string) *CancelResourceExportTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelResourceExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
