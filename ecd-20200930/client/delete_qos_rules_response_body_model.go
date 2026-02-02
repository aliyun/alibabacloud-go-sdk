// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQosRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteQosRulesResponseBody
	GetRequestId() *string
}

type DeleteQosRulesResponseBody struct {
	// example:
	//
	// E54EB497-D7B7-5F04-B744-D8DFA7B******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQosRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQosRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQosRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQosRulesResponseBody) SetRequestId(v string) *DeleteQosRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQosRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
