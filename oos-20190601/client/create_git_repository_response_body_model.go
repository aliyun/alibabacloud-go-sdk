// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGitRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGitRepo(v *CreateGitRepositoryResponseBodyGitRepo) *CreateGitRepositoryResponseBody
	GetGitRepo() *CreateGitRepositoryResponseBodyGitRepo
	SetRequestId(v string) *CreateGitRepositoryResponseBody
	GetRequestId() *string
}

type CreateGitRepositoryResponseBody struct {
	GitRepo *CreateGitRepositoryResponseBodyGitRepo `json:"GitRepo,omitempty" xml:"GitRepo,omitempty" type:"Struct"`
	// example:
	//
	// AA9FA778-AE4B-55EC-81CC-C46BAF08A166
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGitRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGitRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGitRepositoryResponseBody) GetGitRepo() *CreateGitRepositoryResponseBodyGitRepo {
	return s.GitRepo
}

func (s *CreateGitRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGitRepositoryResponseBody) SetGitRepo(v *CreateGitRepositoryResponseBodyGitRepo) *CreateGitRepositoryResponseBody {
	s.GitRepo = v
	return s
}

func (s *CreateGitRepositoryResponseBody) SetRequestId(v string) *CreateGitRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGitRepositoryResponseBody) Validate() error {
	if s.GitRepo != nil {
		if err := s.GitRepo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGitRepositoryResponseBodyGitRepo struct {
	// example:
	//
	// Test parameter for security testing
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// aliyun-computest/lingo
	FullName *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	// example:
	//
	// https://github.com/alibaba/fastjson
	HtmlUrl *string `json:"HtmlUrl,omitempty" xml:"HtmlUrl,omitempty"`
	// example:
	//
	// False
	IsPrivate *string `json:"IsPrivate,omitempty" xml:"IsPrivate,omitempty"`
}

func (s CreateGitRepositoryResponseBodyGitRepo) String() string {
	return dara.Prettify(s)
}

func (s CreateGitRepositoryResponseBodyGitRepo) GoString() string {
	return s.String()
}

func (s *CreateGitRepositoryResponseBodyGitRepo) GetDescription() *string {
	return s.Description
}

func (s *CreateGitRepositoryResponseBodyGitRepo) GetFullName() *string {
	return s.FullName
}

func (s *CreateGitRepositoryResponseBodyGitRepo) GetHtmlUrl() *string {
	return s.HtmlUrl
}

func (s *CreateGitRepositoryResponseBodyGitRepo) GetIsPrivate() *string {
	return s.IsPrivate
}

func (s *CreateGitRepositoryResponseBodyGitRepo) SetDescription(v string) *CreateGitRepositoryResponseBodyGitRepo {
	s.Description = &v
	return s
}

func (s *CreateGitRepositoryResponseBodyGitRepo) SetFullName(v string) *CreateGitRepositoryResponseBodyGitRepo {
	s.FullName = &v
	return s
}

func (s *CreateGitRepositoryResponseBodyGitRepo) SetHtmlUrl(v string) *CreateGitRepositoryResponseBodyGitRepo {
	s.HtmlUrl = &v
	return s
}

func (s *CreateGitRepositoryResponseBodyGitRepo) SetIsPrivate(v string) *CreateGitRepositoryResponseBodyGitRepo {
	s.IsPrivate = &v
	return s
}

func (s *CreateGitRepositoryResponseBodyGitRepo) Validate() error {
	return dara.Validate(s)
}
