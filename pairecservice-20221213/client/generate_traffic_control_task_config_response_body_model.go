// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTrafficControlTaskConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *GenerateTrafficControlTaskConfigResponseBody
	GetConfig() *string
	SetRequestId(v string) *GenerateTrafficControlTaskConfigResponseBody
	GetRequestId() *string
}

type GenerateTrafficControlTaskConfigResponseBody struct {
	Config    *string `json:"Config,omitempty" xml:"Config,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateTrafficControlTaskConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateTrafficControlTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTrafficControlTaskConfigResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GenerateTrafficControlTaskConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateTrafficControlTaskConfigResponseBody) SetConfig(v string) *GenerateTrafficControlTaskConfigResponseBody {
	s.Config = &v
	return s
}

func (s *GenerateTrafficControlTaskConfigResponseBody) SetRequestId(v string) *GenerateTrafficControlTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateTrafficControlTaskConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
