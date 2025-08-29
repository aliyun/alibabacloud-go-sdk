// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneEngineConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngineConfigId(v string) *CloneEngineConfigResponseBody
	GetEngineConfigId() *string
	SetRequestId(v string) *CloneEngineConfigResponseBody
	GetRequestId() *string
}

type CloneEngineConfigResponseBody struct {
	// example:
	//
	// 2
	EngineConfigId *string `json:"EngineConfigId,omitempty" xml:"EngineConfigId,omitempty"`
	// example:
	//
	// A04CB8C0-E74A-5E83-BC61-64D153574EC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneEngineConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneEngineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CloneEngineConfigResponseBody) GetEngineConfigId() *string {
	return s.EngineConfigId
}

func (s *CloneEngineConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneEngineConfigResponseBody) SetEngineConfigId(v string) *CloneEngineConfigResponseBody {
	s.EngineConfigId = &v
	return s
}

func (s *CloneEngineConfigResponseBody) SetRequestId(v string) *CloneEngineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneEngineConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
