// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateExperimentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MigrateExperimentsResponseBody
	GetCode() *string
	SetData(v *MigrateExperimentsResponseBodyData) *MigrateExperimentsResponseBody
	GetData() *MigrateExperimentsResponseBodyData
	SetMessage(v string) *MigrateExperimentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *MigrateExperimentsResponseBody
	GetRequestId() *string
}

type MigrateExperimentsResponseBody struct {
	// example:
	//
	// NO_PERMISSION
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *MigrateExperimentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// NotExistError
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E7C42CC7-2E85-508A-84F4-923B605FD10F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateExperimentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateExperimentsResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateExperimentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *MigrateExperimentsResponseBody) GetData() *MigrateExperimentsResponseBodyData {
	return s.Data
}

func (s *MigrateExperimentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MigrateExperimentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateExperimentsResponseBody) SetCode(v string) *MigrateExperimentsResponseBody {
	s.Code = &v
	return s
}

func (s *MigrateExperimentsResponseBody) SetData(v *MigrateExperimentsResponseBodyData) *MigrateExperimentsResponseBody {
	s.Data = v
	return s
}

func (s *MigrateExperimentsResponseBody) SetMessage(v string) *MigrateExperimentsResponseBody {
	s.Message = &v
	return s
}

func (s *MigrateExperimentsResponseBody) SetRequestId(v string) *MigrateExperimentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateExperimentsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MigrateExperimentsResponseBodyData struct {
	// example:
	//
	// true
	AlreadyExists *bool `json:"AlreadyExists,omitempty" xml:"AlreadyExists,omitempty"`
	// example:
	//
	// draft-8u3ck2or5pw2i4auhf
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// true
	Updated *bool `json:"Updated,omitempty" xml:"Updated,omitempty"`
}

func (s MigrateExperimentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MigrateExperimentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *MigrateExperimentsResponseBodyData) GetAlreadyExists() *bool {
	return s.AlreadyExists
}

func (s *MigrateExperimentsResponseBodyData) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *MigrateExperimentsResponseBodyData) GetUpdated() *bool {
	return s.Updated
}

func (s *MigrateExperimentsResponseBodyData) SetAlreadyExists(v bool) *MigrateExperimentsResponseBodyData {
	s.AlreadyExists = &v
	return s
}

func (s *MigrateExperimentsResponseBodyData) SetExperimentId(v string) *MigrateExperimentsResponseBodyData {
	s.ExperimentId = &v
	return s
}

func (s *MigrateExperimentsResponseBodyData) SetUpdated(v bool) *MigrateExperimentsResponseBodyData {
	s.Updated = &v
	return s
}

func (s *MigrateExperimentsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
