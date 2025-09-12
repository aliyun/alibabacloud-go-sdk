// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSet2Request interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *CreateDataSet2Request
	GetClientId() *int64
	SetName(v string) *CreateDataSet2Request
	GetName() *string
	SetType(v string) *CreateDataSet2Request
	GetType() *string
}

type CreateDataSet2Request struct {
	ClientId *int64  `json:"clientId,omitempty" xml:"clientId,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateDataSet2Request) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSet2Request) GoString() string {
	return s.String()
}

func (s *CreateDataSet2Request) GetClientId() *int64 {
	return s.ClientId
}

func (s *CreateDataSet2Request) GetName() *string {
	return s.Name
}

func (s *CreateDataSet2Request) GetType() *string {
	return s.Type
}

func (s *CreateDataSet2Request) SetClientId(v int64) *CreateDataSet2Request {
	s.ClientId = &v
	return s
}

func (s *CreateDataSet2Request) SetName(v string) *CreateDataSet2Request {
	s.Name = &v
	return s
}

func (s *CreateDataSet2Request) SetType(v string) *CreateDataSet2Request {
	s.Type = &v
	return s
}

func (s *CreateDataSet2Request) Validate() error {
	return dara.Validate(s)
}
