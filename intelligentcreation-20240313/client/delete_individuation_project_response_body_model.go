// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndividuationProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesc(v string) *DeleteIndividuationProjectResponseBody
	GetDesc() *string
	SetRequestId(v string) *DeleteIndividuationProjectResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteIndividuationProjectResponseBody
	GetStatus() *string
}

type DeleteIndividuationProjectResponseBody struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteIndividuationProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndividuationProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndividuationProjectResponseBody) GetDesc() *string {
	return s.Desc
}

func (s *DeleteIndividuationProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIndividuationProjectResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteIndividuationProjectResponseBody) SetDesc(v string) *DeleteIndividuationProjectResponseBody {
	s.Desc = &v
	return s
}

func (s *DeleteIndividuationProjectResponseBody) SetRequestId(v string) *DeleteIndividuationProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndividuationProjectResponseBody) SetStatus(v string) *DeleteIndividuationProjectResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteIndividuationProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
