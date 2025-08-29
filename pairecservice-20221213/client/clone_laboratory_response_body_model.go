// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneLaboratoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLaboratoryId(v string) *CloneLaboratoryResponseBody
	GetLaboratoryId() *string
	SetRequestId(v string) *CloneLaboratoryResponseBody
	GetRequestId() *string
}

type CloneLaboratoryResponseBody struct {
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 01D22D08-BA20-5F35-8302-99115F288220
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneLaboratoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneLaboratoryResponseBody) GoString() string {
	return s.String()
}

func (s *CloneLaboratoryResponseBody) GetLaboratoryId() *string {
	return s.LaboratoryId
}

func (s *CloneLaboratoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneLaboratoryResponseBody) SetLaboratoryId(v string) *CloneLaboratoryResponseBody {
	s.LaboratoryId = &v
	return s
}

func (s *CloneLaboratoryResponseBody) SetRequestId(v string) *CloneLaboratoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneLaboratoryResponseBody) Validate() error {
	return dara.Validate(s)
}
