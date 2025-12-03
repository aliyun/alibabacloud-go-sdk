// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurgeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *PurgeInstanceRequest
	GetClusterId() *string
}

type PurgeInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-m5ek15uzs7613xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s PurgeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s PurgeInstanceRequest) GoString() string {
	return s.String()
}

func (s *PurgeInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *PurgeInstanceRequest) SetClusterId(v string) *PurgeInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *PurgeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
