// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHiveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHiveId(v string) *CreateHiveResponseBody
	GetHiveId() *string
	SetRequestId(v string) *CreateHiveResponseBody
	GetRequestId() *string
}

type CreateHiveResponseBody struct {
	// example:
	//
	// hive-6c1418bf513e400bb697307c077a0ec3
	HiveId *string `json:"HiveId,omitempty" xml:"HiveId,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHiveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHiveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHiveResponseBody) GetHiveId() *string {
	return s.HiveId
}

func (s *CreateHiveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHiveResponseBody) SetHiveId(v string) *CreateHiveResponseBody {
	s.HiveId = &v
	return s
}

func (s *CreateHiveResponseBody) SetRequestId(v string) *CreateHiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHiveResponseBody) Validate() error {
	return dara.Validate(s)
}
