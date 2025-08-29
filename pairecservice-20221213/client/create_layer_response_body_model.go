// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayerId(v string) *CreateLayerResponseBody
	GetLayerId() *string
	SetRequestId(v string) *CreateLayerResponseBody
	GetRequestId() *string
}

type CreateLayerResponseBody struct {
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 59CE7EC6-F268-5D71-9215-32922CC50D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLayerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLayerResponseBody) GetLayerId() *string {
	return s.LayerId
}

func (s *CreateLayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLayerResponseBody) SetLayerId(v string) *CreateLayerResponseBody {
	s.LayerId = &v
	return s
}

func (s *CreateLayerResponseBody) SetRequestId(v string) *CreateLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLayerResponseBody) Validate() error {
	return dara.Validate(s)
}
