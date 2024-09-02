package bundler

import "github.com/DAO-Metaplayer/ai-bundler/pkg/aiop"

func adjustBatchSize(max int, batch []*aiop.AiOperation) []*aiop.AiOperation {
	if len(batch) > max && max > 0 {
		return batch[:max]
	}
	return batch
}
