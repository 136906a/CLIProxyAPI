package management

import "github.com/gin-gonic/gin"

// Quota exceeded toggles
func (h *Handler) GetSwitchProject(c *gin.Context) {
	c.JSON(200, gin.H{"switch-project": h.cfg.QuotaExceeded.SwitchProject})
}
func (h *Handler) PutSwitchProject(c *gin.Context) {
	h.updateBoolField(c, func(v bool) { h.cfg.QuotaExceeded.SwitchProject = v })
}

func (h *Handler) GetSwitchPreviewModel(c *gin.Context) {
	c.JSON(200, gin.H{"switch-preview-model": h.cfg.QuotaExceeded.SwitchPreviewModel})
}
func (h *Handler) PutSwitchPreviewModel(c *gin.Context) {
	h.updateBoolField(c, func(v bool) { h.cfg.QuotaExceeded.SwitchPreviewModel = v })
}

func (h *Handler) GetDisableCredential(c *gin.Context) {
	c.JSON(200, gin.H{"disable-credential": h.cfg.QuotaExceeded.DisableCredential})
}
func (h *Handler) PutDisableCredential(c *gin.Context) {
	h.updateBoolField(c, func(v bool) { h.cfg.QuotaExceeded.DisableCredential = v })
}

func (h *Handler) GetReEnableAfter(c *gin.Context) {
	c.JSON(200, gin.H{"re-enable-after": h.cfg.QuotaExceeded.ReEnableAfter})
}
func (h *Handler) PutReEnableAfter(c *gin.Context) {
	h.updateStringField(c, func(v string) { h.cfg.QuotaExceeded.ReEnableAfter = v })
}
