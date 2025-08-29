// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEngineConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngineConfigId(v string) *CreateEngineConfigResponseBody
	GetEngineConfigId() *string
	SetRequestId(v string) *CreateEngineConfigResponseBody
	GetRequestId() *string
}

type CreateEngineConfigResponseBody struct {
	// example:
	//
	// 1
	EngineConfigId *string `json:"EngineConfigId,omitempty" xml:"EngineConfigId,omitempty"`
	// example:
	//
	// E15A1443-7917-5BE0-AE70-25538ECF398D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEngineConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEngineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEngineConfigResponseBody) GetEngineConfigId() *string {
	return s.EngineConfigId
}

func (s *CreateEngineConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEngineConfigResponseBody) SetEngineConfigId(v string) *CreateEngineConfigResponseBody {
	s.EngineConfigId = &v
	return s
}

func (s *CreateEngineConfigResponseBody) SetRequestId(v string) *CreateEngineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEngineConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
