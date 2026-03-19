// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteAIServiceRequest
	GetDBInstanceId() *string
	SetServiceId(v string) *DeleteAIServiceRequest
	GetServiceId() *string
	SetType(v string) *DeleteAIServiceRequest
	GetType() *string
}

type DeleteAIServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// drama-123456
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// drama
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteAIServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteAIServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteAIServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DeleteAIServiceRequest) GetType() *string {
	return s.Type
}

func (s *DeleteAIServiceRequest) SetDBInstanceId(v string) *DeleteAIServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteAIServiceRequest) SetServiceId(v string) *DeleteAIServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DeleteAIServiceRequest) SetType(v string) *DeleteAIServiceRequest {
	s.Type = &v
	return s
}

func (s *DeleteAIServiceRequest) Validate() error {
	return dara.Validate(s)
}
