// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFpShotDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFpShotDB(v *CreateFpShotDBResponseBodyFpShotDB) *CreateFpShotDBResponseBody
	GetFpShotDB() *CreateFpShotDBResponseBodyFpShotDB
	SetRequestId(v string) *CreateFpShotDBResponseBody
	GetRequestId() *string
}

type CreateFpShotDBResponseBody struct {
	// The details of the media fingerprint library.
	FpShotDB *CreateFpShotDBResponseBodyFpShotDB `json:"FpShotDB,omitempty" xml:"FpShotDB,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFpShotDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFpShotDBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFpShotDBResponseBody) GetFpShotDB() *CreateFpShotDBResponseBodyFpShotDB {
	return s.FpShotDB
}

func (s *CreateFpShotDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFpShotDBResponseBody) SetFpShotDB(v *CreateFpShotDBResponseBodyFpShotDB) *CreateFpShotDBResponseBody {
	s.FpShotDB = v
	return s
}

func (s *CreateFpShotDBResponseBody) SetRequestId(v string) *CreateFpShotDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFpShotDBResponseBody) Validate() error {
	if s.FpShotDB != nil {
		if err := s.FpShotDB.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFpShotDBResponseBodyFpShotDB struct {
	// The configurations of the media fingerprint library.
	//
	// example:
	//
	// null
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The description of the media fingerprint library.
	//
	// example:
	//
	// The library is a text fingerprint library.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the media fingerprint library. We recommend that you keep this ID for subsequent operation calls.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	FpDBId *string `json:"FpDBId,omitempty" xml:"FpDBId,omitempty"`
	// The model ID of the media fingerprint library.
	//
	// example:
	//
	// 11
	ModelId *int32 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The name of the media fingerprint library.
	//
	// example:
	//
	// example-name-****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the media fingerprint library. After the media fingerprint library is created, it enters the **offline*	- state. After the media fingerprint library is processed at the backend, it enters the **active*	- state.
	//
	// example:
	//
	// offline
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateFpShotDBResponseBodyFpShotDB) String() string {
	return dara.Prettify(s)
}

func (s CreateFpShotDBResponseBodyFpShotDB) GoString() string {
	return s.String()
}

func (s *CreateFpShotDBResponseBodyFpShotDB) GetConfig() *string {
	return s.Config
}

func (s *CreateFpShotDBResponseBodyFpShotDB) GetDescription() *string {
	return s.Description
}

func (s *CreateFpShotDBResponseBodyFpShotDB) GetFpDBId() *string {
	return s.FpDBId
}

func (s *CreateFpShotDBResponseBodyFpShotDB) GetModelId() *int32 {
	return s.ModelId
}

func (s *CreateFpShotDBResponseBodyFpShotDB) GetName() *string {
	return s.Name
}

func (s *CreateFpShotDBResponseBodyFpShotDB) GetState() *string {
	return s.State
}

func (s *CreateFpShotDBResponseBodyFpShotDB) SetConfig(v string) *CreateFpShotDBResponseBodyFpShotDB {
	s.Config = &v
	return s
}

func (s *CreateFpShotDBResponseBodyFpShotDB) SetDescription(v string) *CreateFpShotDBResponseBodyFpShotDB {
	s.Description = &v
	return s
}

func (s *CreateFpShotDBResponseBodyFpShotDB) SetFpDBId(v string) *CreateFpShotDBResponseBodyFpShotDB {
	s.FpDBId = &v
	return s
}

func (s *CreateFpShotDBResponseBodyFpShotDB) SetModelId(v int32) *CreateFpShotDBResponseBodyFpShotDB {
	s.ModelId = &v
	return s
}

func (s *CreateFpShotDBResponseBodyFpShotDB) SetName(v string) *CreateFpShotDBResponseBodyFpShotDB {
	s.Name = &v
	return s
}

func (s *CreateFpShotDBResponseBodyFpShotDB) SetState(v string) *CreateFpShotDBResponseBodyFpShotDB {
	s.State = &v
	return s
}

func (s *CreateFpShotDBResponseBodyFpShotDB) Validate() error {
	return dara.Validate(s)
}
