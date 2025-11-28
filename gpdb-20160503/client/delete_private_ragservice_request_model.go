// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateRAGServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeletePrivateRAGServiceRequest
	GetDBInstanceId() *string
}

type DeletePrivateRAGServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DeletePrivateRAGServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateRAGServiceRequest) GoString() string {
	return s.String()
}

func (s *DeletePrivateRAGServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeletePrivateRAGServiceRequest) SetDBInstanceId(v string) *DeletePrivateRAGServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeletePrivateRAGServiceRequest) Validate() error {
	return dara.Validate(s)
}
