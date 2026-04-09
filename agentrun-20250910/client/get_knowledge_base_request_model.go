// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetKnowledgeBaseRequest struct {
}

func (s GetKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
