// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveImageLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type RemoveImageLabelsRequest struct {
}

func (s RemoveImageLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageLabelsRequest) GoString() string {
	return s.String()
}

func (s *RemoveImageLabelsRequest) Validate() error {
	return dara.Validate(s)
}
