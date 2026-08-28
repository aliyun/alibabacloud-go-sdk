// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSkillVersionViaOssRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DownloadSkillVersionViaOssRequest struct {
}

func (s DownloadSkillVersionViaOssRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadSkillVersionViaOssRequest) GoString() string {
	return s.String()
}

func (s *DownloadSkillVersionViaOssRequest) Validate() error {
	return dara.Validate(s)
}
