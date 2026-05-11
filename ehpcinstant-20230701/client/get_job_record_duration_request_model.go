// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobRecordDurationRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetJobRecordDurationRequest struct {
}

func (s GetJobRecordDurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobRecordDurationRequest) GoString() string {
	return s.String()
}

func (s *GetJobRecordDurationRequest) Validate() error {
	return dara.Validate(s)
}
